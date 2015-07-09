package tylerscodebase

import (
	"io/ioutil"
	"mime"
	"net/http"
	"strings"
	"sync"

	"gopkg.in/fsnotify.v1"
)

type response struct {
	mutex    sync.Mutex
	filename string
	mimeType string
	data     []byte
	watcher  *fsnotify.Watcher
}

func newResponse(filename string) (*response, error) {

	var err error
	var resp response

	resp.data, err = ioutil.ReadFile(filename)
	if err != nil {
		return &response{}, err
	}
	resp.filename = filename

	extension := strings.Split(resp.filename, ".")
	if len(extension) > 0 {
		resp.mimeType = mime.TypeByExtension(extension[len(extension)-1])
	} else {
		resp.mimeType = "text/plain"
	}

	resp.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return &response{}, err
	}
	err = resp.watcher.Add(filename)
	if err != nil {
		return &response{}, err
	}
	go newFileUpdater(resp.watcher, &resp.mutex, &resp.data, resp.filename)()

	return &resp, nil
}

func (resp *response) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	resp.mutex.Lock()

	w.Header().Add("Content-Type", resp.mimeType)
	w.Write(resp.data)

	resp.mutex.Unlock()
}
