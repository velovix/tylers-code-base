package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type DocPage struct {
}

type docPageData struct {
	Title string
	Text  template.HTML
}

func NewDocPage() DocPage {

	return DocPage{}
}

func (docPage DocPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Path[len("/docs/"):]

	t, err := template.ParseFiles("./views/documentation.html")
	if err != nil {
		log.Fatalln(err)
	}

	textData, err := ioutil.ReadFile("./public/docs/" + page + ".html")
	if err != nil {
		textData, _ = ioutil.ReadFile("./public/docs/unknown.html")
	}

	docPageData := docPageData{page, template.HTML(string(textData))}
	t.Execute(w, docPageData)
}
