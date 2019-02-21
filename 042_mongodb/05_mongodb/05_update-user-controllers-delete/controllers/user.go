package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jlucktay/golang-web-dev/042_mongodb/05_mongodb/05_update-user-controllers-delete/models"
	"github.com/julienschmidt/httprouter"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type UserController struct {
	users *mongo.Collection
}

func NewUserController(c *mongo.Collection) *UserController {
	return &UserController{c}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	oid, errIDHex := primitive.ObjectIDFromHex(id)

	if errIDHex != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	u := models.User{}

	// Read user
	if err := uc.users.FindOne(context.Background(), oid).Decode(&u); err == nil {
		w.WriteHeader(http.StatusNotFound) // 404
		return
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

	u.Id = primitive.NewObjectID()

	// Create user
	if _, errInsert := uc.users.InsertOne(context.Background(), u); errInsert != nil {
		log.Fatal(errInsert)
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	oid, errIDHex := primitive.ObjectIDFromHex(id)
	if errIDHex != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// Delete user
	dr, errDelete := uc.users.DeleteOne(context.Background(), oid)
	if errDelete != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "Deleted %d user(s): %v\n", dr.DeletedCount, oid)
}
