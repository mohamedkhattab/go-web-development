package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	myhandler := myHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &myhandler,
	}

	server.ListenAndServe()
}
