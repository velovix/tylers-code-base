package main

import (
	"net/http"
)

func main() {

	homepage, err := NewResponse("./views/index.html")
	if err != nil {
		panic(err)
	}

	sd, err := NewStaticDirectory("./public/", "/public/")
	if err != nil {
		panic(err)
	}

	http.Handle("/", homepage)
	http.Handle("/public/", sd)

	http.ListenAndServe(":80", nil)
}
