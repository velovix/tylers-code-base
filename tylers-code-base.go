package tylerscodebase

import (
	"log"
	"net/http"
)

func init() {

	log.SetFlags(log.Ldate | log.Ltime)

	homepage, err := newResponse("./views/index.html")
	if err != nil {
		panic(err)
	}

	favicon, err := newResponse("./public/favicon.ico")
	if err != nil {
		panic(err)
	}

	sd, err := newStaticDirectory("./public/", "/public/")
	if err != nil {
		panic(err)
	}

	docPage := documentation{}

	http.Handle("/", homepage)
	http.Handle("/docs/", docPage)
	http.Handle("/favicon.ico", favicon)
	http.Handle("/public/", sd)
}
