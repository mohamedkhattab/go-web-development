package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// r.Form <- map of form values, ["name"] <- key name, [0] <- number of value in slice
	fmt.Fprintln(w, r.Form["name"][0])
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
