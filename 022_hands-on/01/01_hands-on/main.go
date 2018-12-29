package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "index!")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "doggy!")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "it'sa me, Jimmy!")
}
