package tylerscodebase

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	documentationFile = "./views/documentation.html"
	docDataLocation   = "./public/docs/"
)

type documentation struct {
}

type docPageData struct {
	Title string
	Text  template.HTML
}

func newDocumentation() (*documentation, error) {
	return &documentation{}, nil
}

func (doc *documentation) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Path[len("/docs/"):]

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

	dataWriter := bytes.NewBuffer(make([]byte, 0))
	t.Execute(dataWriter, docPageData)

	// Return results
	w.Write(dataWriter.Bytes())
}
