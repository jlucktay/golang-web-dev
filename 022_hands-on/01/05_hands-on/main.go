package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
	"time"
)

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "index!")
}

type person struct {
	Name string
}

type parseMe struct {
	People []person
	Doggy  string
	Now    time.Time
}

func dog(w http.ResponseWriter, r *http.Request) {
	data := parseMe{
		People: []person{
			{"Evička"},
			{"Karolinka"},
			{"Katka"},
			{"Eva"},
			{"Pawel"},
			{"Jimmy"},
		},
		Doggy: "Rex",
		Now:   time.Now(),
	}
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))

	err := tpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "it'sa me, Jimmy!")
}
