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

	doc, err := newDocumentation()
	if err != nil {
		panic(err)
	}

	http.Handle("/", homepage)
	http.Handle("/docs/", doc)
	http.Handle("/favicon.ico", favicon)
}
