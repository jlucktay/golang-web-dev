package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogJpg)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := template.Must(template.ParseFiles("dog.gohtml")).Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Served '/dog/'! (%s)\n", time.Now())
}

func dogJpg(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")

	fmt.Printf("Served '/dog.jpg'! (%s)\n", time.Now())
}
