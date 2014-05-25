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

	favicon, icoErr := NewResponse("./public/favicon.ico")
	if err != nil {
		panic(icoErr)
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
