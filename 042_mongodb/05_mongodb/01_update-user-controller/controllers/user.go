package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jlucktay/golang-web-dev/042_mongodb/05_mongodb/01_update-user-controller/models"
	"github.com/julienschmidt/httprouter"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// added session to our userController
type UserController struct {
	client *mongo.Client
}

// added session to our userController
func NewUserController(c *mongo.Client) *UserController {
	return &UserController{c}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

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

	u.Id = "007"

	uj, errMarshal := json.Marshal(u)
	if errMarshal != nil {
		log.Fatal(errMarshal)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: only write code to delete user
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}
