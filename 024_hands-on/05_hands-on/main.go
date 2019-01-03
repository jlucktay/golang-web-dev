package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	err := template.Must(template.ParseFiles("templates/index.gohtml")).Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Served '/index.html'! (%s)\n", time.Now())
}
