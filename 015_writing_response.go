package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Thread struct {
	Name   string
	Number int64
}

func response(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><tutle>Response Written</title></head>
<body><h1>Response is here</h1></body>
</html>`
	w.Write([]byte(str))
}

func header(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(403)
	fmt.Fprintln(w, "403 Access to this resource is forbidden")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.google.com/")
	w.WriteHeader(302)
}

func returnJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	thread := &Thread{
		Name:   "My First Thread",
		Number: 123,
	}
	json, _ := json.Marshal(thread)
	w.Write(json)
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/response", response)
	http.HandleFunc("/writeheader", header)
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/json", returnJSON)

	server.ListenAndServe()
}
