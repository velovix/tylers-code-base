package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	DefaultViewCount = 0
	DefaultPort      = 3000
)

var (
	viewCount int
	port      int
)

func loadConfigFile() {

	file, err := os.Open("./config.json")
	if err != nil {
		newConfigFile()
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
	port = int(config["port"].(float64))
}

func saveConfigFile() {

	file, err := os.Create("./config.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)

	config := make(map[string]interface{})
	config["viewCount"] = viewCount
	config["port"] = port

	err = encoder.Encode(config)
	if err != nil {
		panic(err)
	}
}

func newConfigFile() {

	file, err := os.Create("./config.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)

	config := make(map[string]interface{})
	config["viewCount"] = DefaultViewCount
	config["port"] = DefaultPort

	err = encoder.Encode(config)
	if err != nil {
		panic(err)
	}
}

func addView(w http.ResponseWriter, r *http.Request) {

	viewCount++
	fmt.Fprintf(w, "%v", viewCount)
	saveConfigFile()
}

func main() {

	log.SetFlags(log.Ldate | log.Ltime)

	loadConfigFile()

	homepage, err := NewResponse("./views/index.html")
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

	docPage := DocPage{}

	http.HandleFunc("/addView", addView)
	http.Handle("/", homepage)
	http.Handle("/docs/", docPage)
	http.Handle("/favicon.ico", favicon)
	http.Handle("/public/", sd)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
