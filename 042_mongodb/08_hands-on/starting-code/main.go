package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jlucktay/golang-web-dev/042_mongodb/08_hands-on/starting-code/controllers"
	"github.com/jlucktay/golang-web-dev/042_mongodb/08_hands-on/starting-code/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func getSession() (m map[string]models.User) {
	m = make(map[string]models.User)

	if _, errExist := os.Stat(controllers.PersistentStorage); os.IsNotExist(errExist) {
		if _, errCreate := os.Create(controllers.PersistentStorage); errCreate != nil {
			log.Fatal(errCreate)
		}
		return
	}

	ps, errOpen := os.Open(controllers.PersistentStorage)
	if errOpen != nil {
		log.Fatal(errOpen)
	}
	defer ps.Close()

	if errDecode := json.NewDecoder(ps).Decode(&m); errDecode != nil {
		log.Fatal(errDecode)
	}

	return
}
