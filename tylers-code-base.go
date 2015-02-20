package tylerscodebase

import (
	"log"
	"net/http"
)

func init() {

	log.SetFlags(log.Ldate | log.Ltime)

	homepage, err := NewResponse("./views/index.html")
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

	docPage := DocPage{}

	http.Handle("/", homepage)
	http.Handle("/docs/", docPage)
	http.Handle("/favicon.ico", favicon)
	http.Handle("/public/", sd)
}
