package main

import (
	"net/http"

	"github.com/jlucktay/golang-web-dev/042_mongodb/05_mongodb/05_update-user-controllers-delete/controllers"
	"github.com/julienschmidt/httprouter"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mongo.Session {
	s, err := mongo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
