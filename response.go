package tylerscodebase

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type response struct {
	file string
}

func newResponse(filename string) (response, error) {

	_, err := ioutil.ReadFile(filename)
	if err != nil {
		return response{""}, err
	}

	return response{filename}, nil
}

func (resp response) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open(resp.file)
	if err != nil {
		fmt.Fprint(w, "404 file error: %s", resp.file)
		return
	}

	var zeroTime time.Time
	http.ServeContent(w, r, resp.file, zeroTime, file)
}
