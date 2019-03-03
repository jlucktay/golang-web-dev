package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jlucktay/golang-web-dev/042_mongodb/05_mongodb/05_update-user-controllers-delete/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	// Read user
	u := models.User{}
	readUser := uc.users.FindOne(
		context.Background(),
		primitive.D{
			{Key: "_id", Value: oid},
		},
	)
	if errDecode := readUser.Decode(&u); errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, "error decoding user: %s\n", errDecode)
		return
	}
	mu, errMarshal := json.MarshalIndent(u, "", "  ")
	if errMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, "error marshaling user: %s\n", errMarshal)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", mu)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	if errDecode := json.NewDecoder(r.Body).Decode(&u); errDecode != nil {
		log.Fatal(errDecode)
	}

	u.Id = primitive.NewObjectID()

	// Create user in DB
	insertUser := primitive.D{
		{Key: "_id", Value: u.Id},
		{Key: "age", Value: u.Age},
		{Key: "gender", Value: u.Gender},
		{Key: "name", Value: u.Name},
	}
	ior, errInsert := uc.users.InsertOne(context.Background(), insertUser)
	if errInsert != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, "error inserting user: %s\n", errInsert)
		return
	}

	// Return result
	mu, errMarshal := json.MarshalIndent(u, "", "  ")
	if errMarshal != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, "error marshaling user: %s\n", errMarshal)
		return
	}

	fmt.Printf("insert result: %+v\n", ior)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", mu)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	oid, errIDHex := primitive.ObjectIDFromHex(id)
	if errIDHex != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// Delete user
	deleteUser := primitive.D{
		{Key: "_id", Value: oid},
	}
	dr, errDelete := uc.users.DeleteOne(context.Background(), deleteUser)
	if errDelete != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	fmt.Printf("delete result: %+v\n", dr)
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "Deleted %d user(s): %v\n", dr.DeletedCount, oid)
}

func (uc UserController) ResetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	count, errCount := uc.users.CountDocuments(context.Background(), primitive.D{})
	if errCount != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, "Error counting users: %v\n", errCount)
		return
	}

	if errDrop := uc.users.Drop(context.Background()); errDrop != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, "Error dropping users: %v\n", errDrop)
		return
	}

	result := fmt.Sprintf("Dropped collection containing %d users.\n", count)
	fmt.Print(result)
	w.Header().Add(http.CanonicalHeaderKey("drop-count"), strconv.FormatInt(count, 10))
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, result)
}
