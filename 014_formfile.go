package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func printFile(w http.ResponseWriter, r *http.Request) {
	// r.FormFile ideal for a form with a single file
	file, _, err := r.FormFile("uploaded")
	if err != nil {
		fmt.Println(err)
	} else {
		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/printfile", printFile)

	server.ListenAndServe()
}
