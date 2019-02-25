package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/jlucktay/golang-web-dev/042_mongodb/10_hands-on/starting-code/models"
	"github.com/jlucktay/golang-web-dev/042_mongodb/10_hands-on/starting-code/session"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func (c Controller) SignUp(w http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username taken?
		if _, ok := session.DbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		ck := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		ck.MaxAge = session.SessionLength
		http.SetCookie(w, ck)
		session.DbSessions[ck.Value] = models.Session{UN: un, LastActivity: time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = models.User{
			UserName: un,
			Password: bs,
			First:    f,
			Last:     l,
			Role:     r,
		}
		session.DbUsers[un] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	session.ShowSessions() // for demonstration purposes
	err := c.tpl.ExecuteTemplate(w, "signup.gohtml", u)
	if err != nil {
		log.Fatal(err)
	}

}

func (c Controller) Login(w http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := session.DbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		ck := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		ck.MaxAge = session.SessionLength
		http.SetCookie(w, ck)
		session.DbSessions[ck.Value] = models.Session{UN: un, LastActivity: time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	session.ShowSessions() // for demonstration purposes
	err := c.tpl.ExecuteTemplate(w, "login.gohtml", u)
	if err != nil {
		log.Fatal(err)
	}

}

func (c Controller) Logout(w http.ResponseWriter, req *http.Request) {
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	ck, _ := req.Cookie("session")
	// delete the session
	delete(session.DbSessions, ck.Value)
	// remove the cookie
	ck = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, ck)

	// clean up dbSessions
	if time.Since(session.DbSessionsCleaned) > (time.Second * 30) {
		go session.CleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
