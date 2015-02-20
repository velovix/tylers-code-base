package tylerscodebase

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Response struct {
	file string
}

func NewResponse(filename string) (Response, error) {

	_, err := ioutil.ReadFile(filename)
	if err != nil {
		return Response{""}, err
	}

	return Response{filename}, nil
}

func (response Response) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open(response.file)
	if err != nil {
		fmt.Fprint(w, "404 file error: %s", response.file)
		return
	}

	var zeroTime time.Time
	http.ServeContent(w, r, response.file, zeroTime, file)
}
