package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jlucktay/golang-web-dev/042_mongodb/03_post-delete/models"
	"github.com/julienschmidt/httprouter"
)

var users map[string]models.User

func init() {
	users = make(map[string]models.User)
}

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	// added route
	r.POST("/user", createUser)
	// added route plus parameter
	r.DELETE("/user/:id", deleteUser)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Index</title>
</head>
<body>
<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
</body>
</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(s)); err != nil {
		log.Fatal(err)
	}
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	existingUser, exists := users[p.ByName("id")]

	if !exists {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "user ID '%s' not found\n", p.ByName("id"))
		return
	}

	// Marshal into JSON
	uj, err := json.Marshal(existingUser)
	if err != nil {
		fmt.Println(err)
	}

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// composite literal - type and curly braces
	u := models.User{}

	// encode/decode for sending/receiving JSON to/from a stream
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Fatal(err)
	}

	if _, exists := users[u.Id]; exists {
		w.WriteHeader(http.StatusConflict) // 409
		fmt.Fprint(w, "user already exists\n")
		return
	}

	users[u.Id] = u

	// marshal/unmarshal for having JSON assigned to a variable
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if _, exists := users[p.ByName("id")]; !exists {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "user ID '%s' not found\n", p.ByName("id"))
		return
	}

	delete(users, p.ByName("id"))

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "user ID '%s' deleted\n", p.ByName("id"))
}
