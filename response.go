package tylerscodebase

import (
	"io/ioutil"
	"mime"
	"net/http"
	"strings"
)

type response struct {
	filename string
}

func newResponse(filename string) (*response, error) {

	var resp response

	resp.filename = filename

	return &resp, nil
}

func (resp *response) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile(resp.filename)
	if err != nil {
		http.Error(w, "could not read file", 500)
		return
	}

	var mimeType string
	extension := strings.Split(resp.filename, ".")
	if len(extension) > 0 {
		mimeType = mime.TypeByExtension(extension[len(extension)-1])
	} else {
		mimeType = "text/plain"
	}
	w.Header().Add("Content-Type", mimeType)

	w.Write(data)
}
