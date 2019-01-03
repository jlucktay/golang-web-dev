package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", dogs)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatal("template didn't execute: ", err)
	}

	fmt.Printf("Served '/'! (%s)\n", time.Now())
}
