package tylerscodebase

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"gopkg.in/fsnotify.v1"
)

// newFileUpdater returns a function that waits for a file modification and
// updates the data slice.
func newFileUpdater(watcher *fsnotify.Watcher, mutex *sync.Mutex, data *[]byte, filename string) func() {

	return func() {
		var err error

		for {
			select {
			case event := <-watcher.Events:
				mutex.Lock()
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", filename)
					*data, err = ioutil.ReadFile(filename)
					if err != nil {
						log.Println(err)
					}
				}
				mutex.Unlock()
			case err := <-watcher.Errors:
				log.Println(err)
			}
		}
	}
}

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
