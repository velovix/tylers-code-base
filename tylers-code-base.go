package main

import (
	"net/http"
)

func main() {

	homepage, err := NewResponse("./views/index.html")
	if err != nil {
		panic(err)
	}

	http.Handle("/", homepage)

	http.ListenAndServe(":80", nil)
}
