package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

	data, err := ioutil.ReadFile(response.file)
	if err != nil {
		fmt.Fprint(w, "404 file error: %s", response.file)
	}

	fmt.Fprint(w, string(data))
}
