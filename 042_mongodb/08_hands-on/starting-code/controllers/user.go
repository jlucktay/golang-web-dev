package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jlucktay/golang-web-dev/042_mongodb/08_hands-on/starting-code/models"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

const (
	PersistentStorage = "./persistent.json"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Retrieve user
	u := uc.session[id]

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	if errDecode := json.NewDecoder(r.Body).Decode(&u); errDecode != nil {
		log.Fatal(errDecode)
	}

	// create ID
	newId, _ := uuid.NewV4()
	u.Id = newId.String()

	// store the user
	uc.session[u.Id] = u

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	uc.saveSession()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.session, id)

	uc.saveSession()
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "Deleted user '%s'.\n", id)
}

func (uc UserController) saveSession() {
	ps, errOpen := os.OpenFile(PersistentStorage, os.O_WRONLY|os.O_TRUNC, 0644)
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	enc := json.NewEncoder(ps)
	if errEncode := enc.Encode(uc.session); errEncode != nil {
		log.Fatal(errEncode)
	}
}
