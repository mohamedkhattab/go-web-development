package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct{}

func (mh *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

type worldHandler struct{}

func (wh *worldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {

	whandler := worldHandler{}
	hhandler := helloHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello", &hhandler)
	http.Handle("/world", &whandler)

	server.ListenAndServe()
}
