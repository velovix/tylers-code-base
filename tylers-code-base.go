package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	viewCount int
)

func addView(w http.ResponseWriter, r *http.Request) {

	viewCount++
	fmt.Fprintf(w, "%v", viewCount)
	saveConfigFile()
}

func loadConfigFile() {

	file, err := os.Open("./config.json")
	if err != nil {
		saveConfigFile()
		loadConfigFile()
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	config := make(map[string]interface{})
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	viewCount = int(config["viewCount"].(float64))
}

func saveConfigFile() {

	file, err := os.Create("./config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)

	config := make(map[string]interface{})
	config["viewCount"] = viewCount

	err = encoder.Encode(config)
	if err != nil {
		panic(err)
	}
}

func main() {

	loadConfigFile()

	homepage, err := NewResponse("./views/index.html")
	if err != nil {
		panic(err)
	}

	documentation, err := NewResponse("./views/documentation.html")
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

	http.HandleFunc("/addView", addView)
	http.Handle("/", homepage)
	http.Handle("/docs", documentation)
	http.Handle("/favicon.ico", favicon)
	http.Handle("/public/", sd)

	http.ListenAndServe(":3000", nil)
}
