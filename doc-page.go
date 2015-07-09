package tylerscodebase

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"gopkg.in/fsnotify.v1"
)

const documentationFile = "./views/documentation.html"

type documentation struct {
	mutex        sync.Mutex
	pageData     map[string]*[]byte
	pageWatchers map[string]*fsnotify.Watcher
}

type docPageData struct {
	Title string
	Text  template.HTML
}

func newDocumentation() (*documentation, error) {

	var doc documentation

	doc.pageData = make(map[string]*[]byte)
	doc.pageWatchers = make(map[string]*fsnotify.Watcher)

	return &doc, nil
}

func (doc *documentation) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Path[len("/docs/"):]

	// Check if the page is already cached
	if data, ok := doc.pageData[page]; ok {
		w.Write(*data)
	} else {
		doc.mutex.Lock()

		// Parse the generic documentation template
		t, err := template.ParseFiles("./views/documentation.html")
		if err != nil {
			log.Fatalln(err)
		}

		// Read the documentation data file
		textData, err := ioutil.ReadFile("./public/docs/" + page + ".html")
		if err != nil {
			textData, _ = ioutil.ReadFile("./public/docs/unknown.html")
			page = "unknown"
		}

		// Create a struct to fill the template with
		docPageData := docPageData{page, template.HTML(string(textData))}

		// Save the completed page into the cache
		doc.pageData[page] = new([]byte)
		dataWriter := bytes.NewBuffer(make([]byte, 0))
		t.Execute(dataWriter, docPageData)
		*doc.pageData[page] = dataWriter.Bytes()

		// Start watching the data file
		doc.pageWatchers[page], err = fsnotify.NewWatcher()
		if err != nil {
			log.Println(err)
			return
		}
		err = doc.pageWatchers[page].Add("./public/docs/" + page + ".html")
		go newFileUpdater(doc.pageWatchers[page], &doc.mutex, doc.pageData[page], "./public/docs/"+page+".html")()

		// Return results
		w.Write(*doc.pageData[page])

		doc.mutex.Unlock()
	}
}
