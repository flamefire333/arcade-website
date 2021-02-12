package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

var count int = 0

func request(incoming chan int, outgoing chan int) {
	for 1 == 1 {
		requestType := <-incoming
		if requestType == 0 {
			count = count + 1
		} else if requestType == 1 {
			outgoing <- count
		}
	}
}

func main() {
	r := mux.NewRouter()
	sendChannel := make(chan int)
	receiveChannel := make(chan int)
	go request(sendChannel, receiveChannel)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>This is the homepage</h1>")
		sendChannel <- 0
	})

	r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["name"]
		sendChannel <- 1
		visit := <-receiveChannel
		fmt.Fprintf(w, "<h1>Hello, %s! This is visit %d\n</h1>", title, visit)
	})

	http.ListenAndServe(":80", r)
}
