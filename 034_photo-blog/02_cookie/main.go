package main

import (
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

const cookieName = "session"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	if err := tpl.ExecuteTemplate(w, "index.gohtml", c.Value); err != nil {
		log.Fatal(err)
	}
}

// add func to get cookie
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie(cookieName)
	if err != nil {
		sID := uuid.Must(uuid.NewV4())
		c = &http.Cookie{
			Name:  cookieName,
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}
