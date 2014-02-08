package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type StaticDirectory struct {
	path    string
	trigger string
}

func NewStaticDirectory(path, trigger string) (StaticDirectory, error) {

	_, err := ioutil.ReadDir(path)
	if err != nil {
		return StaticDirectory{"", ""}, err
	}

	return StaticDirectory{path, trigger}, nil
}

func (staticDirectory StaticDirectory) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	trimmedDir := staticDirectory.path + r.URL.Path[len(staticDirectory.trigger):]

	file, err := os.Open(trimmedDir)
	if err != nil {
		fmt.Fprint(w, "404 file error: %s", trimmedDir)
		return
	}

	var zeroTime time.Time
	http.ServeContent(w, r, trimmedDir, zeroTime, file)
}
