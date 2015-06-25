package tylerscodebase

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type staticDirectory struct {
	path    string
	trigger string
}

func newStaticDirectory(path, trigger string) (staticDirectory, error) {

	_, err := ioutil.ReadDir(path)
	if err != nil {
		return staticDirectory{"", ""}, err
	}

	return staticDirectory{path, trigger}, nil
}

func (sd staticDirectory) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	trimmedDir := sd.path + r.URL.Path[len(sd.trigger):]

	file, err := os.Open(trimmedDir)
	if err != nil {
		fmt.Fprint(w, "404 file error: %s", trimmedDir)
		return
	}

	var zeroTime time.Time
	http.ServeContent(w, r, trimmedDir, zeroTime, file)
}
