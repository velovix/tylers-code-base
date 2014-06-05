package main

import (
	"net/http"
)

func main() {

	homepage, err := NewResponse("./views/index.html")
	if err != nil {
		panic(err)
	}

	documentation, err := NewResponse("./views/documentation.html")
	if err != nil {
		panic(err)
	}

	favicon, err := NewResponse("./public/favicon.ico")
	if err != nil {
		panic(err)
	}

	sd, err := NewStaticDirectory("./public/", "/public/")
	if err != nil {
		panic(err)
	}

	http.Handle("/", homepage)
	http.Handle("/docs", documentation)
	http.Handle("/favicon.ico", favicon)
	http.Handle("/public/", sd)

	http.ListenAndServe(":3000", nil)
}
