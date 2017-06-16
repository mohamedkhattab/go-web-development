package main

import (
	"fmt"
	"net/http"
)

func sendCookies(w http.ResponseWriter, r *http.Request) {
	cookie1 := http.Cookie{
		Name:     "username",
		Value:    "BernieSanders2020",
		HttpOnly: true,
	}

	cookie2 := http.Cookie{
		Name:     "password",
		Value:    "Password123",
		HttpOnly: true,
	}

	w.Header().Set("Set-Cookie", cookie1.String())
	w.Header().Add("Set-Cookie", cookie2.String())
}

func sendCookies2(w http.ResponseWriter, r *http.Request) {
	cookie1 := http.Cookie{
		Name:     "username",
		Value:    "RonPaul2020",
		HttpOnly: true,
	}

	cookie2 := http.Cookie{
		Name:     "password",
		Value:    "Password777",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Header["Cookie"])
}

func getCookie2(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the cookie")
	}
	fmt.Fprintln(w, username.Value)

	cs := r.Cookies()
	fmt.Fprintln(w, cs)
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/sendcookies", sendCookies)
	http.HandleFunc("/sendcookies2", sendCookies2)
	http.HandleFunc("/getcookie", getCookie)
	http.HandleFunc("/getcookie2", getCookie2)

	server.ListenAndServe()
}
