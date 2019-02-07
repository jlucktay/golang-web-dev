package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

const cookieName = "my-session-uuid"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	if err := tpl.ExecuteTemplate(w, "index.gohtml", c.Value); err != nil {
		log.Fatal(err)
	}
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie(cookieName)
	if err != nil {
		u := uuid.Must(uuid.NewV4())
		c := &http.Cookie{
			Name:  cookieName,
			Value: u.String(),
		}
		http.SetCookie(w, c)
	}

	return c
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie(cookieName)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Your UUID cookie:", c)
}
