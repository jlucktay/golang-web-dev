package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jlucktay/golang-web-dev/042_mongodb/10_hands-on/starting-code/models"
	uuid "github.com/satori/go.uuid"
)

const SessionLength int = 30

var (
	DbSessions        = map[string]models.Session{} // session ID, session
	DbSessionsCleaned time.Time
	DbUsers           = map[string]models.User{} // user ID, user
)

func init() {
	DbSessionsCleaned = time.Now()
}

func GetUser(w http.ResponseWriter, req *http.Request) models.User {
	// get cookie
	ck, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		ck = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	ck.MaxAge = SessionLength
	http.SetCookie(w, ck)

	// if the user exists already, get user
	var u models.User
	if s, ok := DbSessions[ck.Value]; ok {
		s.LastActivity = time.Now()
		DbSessions[ck.Value] = s
		u = DbUsers[s.UserName]
	}
	return u
}
func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	ck, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := DbSessions[ck.Value]
	if ok {
		s.LastActivity = time.Now()
		DbSessions[ck.Value] = s
	}
	_, ok = DbUsers[s.UserName]
	// refresh session
	ck.MaxAge = SessionLength
	http.SetCookie(w, ck)
	return ok
}

func CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	ShowSessions()              // for demonstration purposes
	for k, v := range DbSessions {
		if time.Since(v.LastActivity) > (time.Second * 30) {
			delete(DbSessions, k)
		}
	}
	DbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	ShowSessions()             // for demonstration purposes
}

// for demonstration purposes
func ShowSessions() {
	fmt.Println("********")
	for k, v := range DbSessions {
		fmt.Println(k, v.UserName)
	}
	fmt.Println("")
}
