package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func sendMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Message was set successfully!")
	c1 := http.Cookie{
		Name:  "flash_msg",
		Value: base64.URLEncoding.EncodeToString(msg),
	}

	http.SetCookie(w, &c1)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash_msg")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Println(w, "Cookie wasn't found")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash_msg",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}

		http.SetCookie(w, &rc)

		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/sendmessage", sendMessage)
	http.HandleFunc("/showmessage", showMessage)

	server.ListenAndServe()
}
