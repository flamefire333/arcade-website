package main

import (
	"arcade-website/chat"
	"arcade-website/mafia"
	"arcade-website/roles"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Accessing Path %s", path)

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		log.Printf("%s does not exist.\n", path)
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Serving %s from static.\n", path)
	// otherwise, use http.FileServer to serve the static dir
	// http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
	http.ServeFile(w, r, path)
}

func main() {
	r := mux.NewRouter()
	rand.Seed(time.Now().UnixNano())
	roles.InitRoles()
	go chat.ChatRequestHandler()
	go mafia.MafiaRequestHandler()
	go mafia.GameRunner()

	chat.AddChatHandlers(r)
	mafia.AddMafiaHandlers(r)

	spa := spaHandler{staticPath: "/frontend", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	http.ListenAndServe(":80", r)
}
