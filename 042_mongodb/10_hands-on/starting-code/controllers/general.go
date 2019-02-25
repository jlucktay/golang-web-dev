package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jlucktay/golang-web-dev/042_mongodb/10_hands-on/starting-code/session"
)

type Controller struct {
	tpl *template.Template
}

func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

func (c Controller) Index(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	session.ShowSessions() // for demonstration purposes
	err := c.tpl.ExecuteTemplate(w, "index.gohtml", u)
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) Bar(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	session.ShowSessions() // for demonstration purposes
	err := c.tpl.ExecuteTemplate(w, "bar.gohtml", u)
	if err != nil {
		log.Fatal(err)
	}

}
