package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/jlucktay/golang-web-dev/042_mongodb/06_hands-on/starting-code/models"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	users map[string]models.User
}

func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// User IDs are 16-digit hexadecimal strings
	idRx := regexp.MustCompile(`^[0-9a-f]{16}$`)
	if !idRx.Match([]byte(id)) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// Fetch user
	existing, exists := uc.users[id]
	if !exists {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(existing)

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
	u.Id = RandStringBytesMaskImprSrc(16)

	// store the user
	uc.users[u.Id] = u

	uj, errMarshal := json.Marshal(u)
	if errMarshal != nil {
		log.Fatal(errMarshal)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// User IDs are 16-digit hexadecimal strings
	idRx := regexp.MustCompile(`^[0-9a-f]{16}$`)
	if !idRx.Match([]byte(id)) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	delete(uc.users, id)
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "Deleted user '%s'\n", id)
}
