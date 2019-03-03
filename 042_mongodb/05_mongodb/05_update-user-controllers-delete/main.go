package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jlucktay/golang-web-dev/042_mongodb/05_mongodb/05_update-user-controllers-delete/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Printf("[%v] Here we go...\n", time.Now().Format(time.RFC3339))

	r := httprouter.New()
	uc := controllers.NewUserController(getUserCollection())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	r.HEAD("/user/reset", uc.ResetUsers)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func getUserCollection() *mongo.Collection {
	fmt.Printf("[%v] Connecting to MongoDB...\n", time.Now().Format(time.RFC3339))

	// Connect to our local mongo
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	client, errConnect := mongo.Connect(ctx, "mongodb://localhost:27017")

	// Check if connection error, is mongo running?
	if errConnect != nil {
		panic(errConnect)
	}

	fmt.Printf("[%v] Pinging MongoDB connection...\n", time.Now().Format(time.RFC3339))

	if errPing := client.Ping(ctx, readpref.Primary()); errPing != nil {
		panic(errPing)
	}

	fmt.Printf("[%v] Looks like we have successfully connected to MongoDB!\n", time.Now().Format(time.RFC3339))

	collection := client.Database("042_mongodb").Collection("05_users")

	return collection
}
