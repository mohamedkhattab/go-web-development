package main

import (
	"fmt"
	"net/http"
)

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Header["Accept-Language"])
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/headers", header)

	server.ListenAndServe()
}
