package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct{}

func (hh helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Log function called - main.log")
		h.ServeHTTP(w, r)
	})
}

func secure(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Secure Function called - main.secure")
		h.ServeHTTP(w, r)
	})
}

func main() {

	hhanlder := helloHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello", secure(log(hhanlder)))

	server.ListenAndServe()
}
