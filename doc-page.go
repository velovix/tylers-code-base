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

const (
	documentationFile = "./views/documentation.html"
	docDataLocation   = "./public/docs/"
)

type documentation struct {
	mutex        sync.Mutex
	watcher      *fsnotify.Watcher
	pageData     map[string]*[]byte
	pageWatchers map[string]*fsnotify.Watcher
}

type docPageData struct {
	Title string
	Text  template.HTML
}

func newDocumentation() (*documentation, error) {

	var err error
	var doc documentation

	// Start watching the doc template file
	doc.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return &documentation{}, err
	}
	err = doc.watcher.Add(documentationFile)
	if err != nil {
		return &documentation{}, err
	}
	go func() {
		for {
			select {
			case event := <-doc.watcher.Events:
				doc.mutex.Lock()
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", documentationFile)
					// Clear the cache if the doc template has been modified
					doc.pageData = make(map[string]*[]byte)
				}
				doc.mutex.Unlock()
			case err := <-doc.watcher.Errors:
				log.Println(err)
			}
		}
	}()

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
		t, err := template.ParseFiles(documentationFile)
		if err != nil {
			log.Fatalln(err)
		}

		// Read the documentation data file
		textData, err := ioutil.ReadFile(docDataLocation + page + ".html")
		if err != nil {
			textData, _ = ioutil.ReadFile(docDataLocation + "unknown.html")
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
		if err != nil {
			log.Println(err)
			return
		}
		go newFileUpdater(doc.pageWatchers[page], &doc.mutex, doc.pageData[page], "./public/docs/"+page+".html")()

		// Return results
		w.Write(*doc.pageData[page])

		doc.mutex.Unlock()
	}
}
