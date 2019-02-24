package main

import (
	"log"
	"net/http"

	"github.com/jlucktay/golang-web-dev/042_mongodb/06_hands-on/starting-code/controllers"
	"github.com/jlucktay/golang-web-dev/042_mongodb/06_hands-on/starting-code/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getMap())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func getMap() map[string]models.User {
	m := make(map[string]models.User)

	return m
}
