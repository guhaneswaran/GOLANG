package main

import (
	"../dao"
	"../models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

var ash = Trainer{"Ash", 10, "Pallet Town"}

func errorLog(e1 error) {
	if e1 != nil {
		log.Fatal(e1)
	}
}

func init() {

	//Set Client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	errorLog(err)

	//Check the connection
	err = client.Ping(context.TODO(), nil)

	errorLog(err)

	fmt.Println("Connected to MONGODB!!")

	// Get a handle for your collection
	collection = client.Database("test").Collection("trainers")

}

func main() {

	//Create a new router
	router := mux.NewRouter()

	//Router Handlers
	router.HandleFunc("/api/trainer", createTrainer).Methods("POST")
	router.HandleFunc("/api/trainer", getAllTrainers).Methods("GET")
	router.HandleFunc("/api/trainer/{name}", getTrainer).Methods("GET")
	router.HandleFunc("/api/trainer/{name}", deleteTrainer).Methods("DELETE")
	router.HandleFunc("/api/trainer/{name}", updateTrainer).Methods("PUT")

	fmt.Println("HTTP listening on http://localhost:9000")
	//Listen
	log.Fatal(http.ListenAndServe(":9000", router))

	// Some dummy data to add to the Database

	// misty := Trainer{"Misty", 10, "Cerulean City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}
	// gary := Trainer{"Gary", 10, "Pallet Town"}

	// // Insert multiple documents
	// trainers := []interface{}{misty, brock, gary}

	// insertIDs := createTrainers(collection, trainers)
	// fmt.Println("Inserted multiple documents: ", insertIDs)

}
