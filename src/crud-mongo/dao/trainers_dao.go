package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"../../../D:/Workspaces/GO LANG/src/crud-mongo/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createTrainer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var trainer Trainer
	_ = json.NewDecoder(r.Body).Decode(&trainer)

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), trainer)
	errorLog(err)

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}

func createTrainers(collection *mongo.Collection, trainers []interface{}) interface{} {

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	errorLog(err)

	return insertManyResult.InsertedIDs
}

func getAllTrainers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	findOptions := options.Find()
	//findOptions.SetLimit(2)

	var trainers []*Trainer
	curs, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	errorLog(err)

	for curs.Next(context.TODO()) {

		//Create a value in which a single document can be decoded

		var singleTrainer Trainer
		err := curs.Decode(&singleTrainer)

		errorLog(err)

		trainers = append(trainers, &singleTrainer)

	}

	for _, iterResults := range trainers {

		fmt.Printf("Items found %+v\n", *iterResults)
	}

	_ = json.NewEncoder(w).Encode(trainers)

}

func getTrainer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	filter := bson.D{{"name", params["name"]}}

	var result Trainer

	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	errorLog(err)

	_ = json.NewEncoder(w).Encode(result)

}

func deleteTrainer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	filter := bson.D{{"name", params["name"]}}

	deleteResult, err := collection.DeleteMany(context.TODO(), filter)

	errorLog(err)

	fmt.Println("Delete entry - ", deleteResult.DeletedCount)

	_ = json.RawMessage("Deleted entries successfully")

}

func updateTrainer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var trainer Trainer

	err := json.NewDecoder(r.Body).Decode(&trainer)

	filter := bson.D{{"name", params["name"]}}

	update := bson.D{
		{"$set", bson.D{
			{"name", trainer.Name},
			{"age", trainer.Age},
			{"city", trainer.City},
		},
		},
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update)

	errorLog(err)

}
