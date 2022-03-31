package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"productlist/app/handler"
	"productlist/app/model"
	"productlist/app/schema"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreatePebbles will handle the create pebbles post request
func CreatePebbles(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	pebbles := new(schema.Pebbles)
	err := json.NewDecoder(req.Body).Decode(pebbles)
	if err != nil {
		handler.ResponseWriter(res, http.StatusBadRequest, "body json request have issues!!!", nil)
		return
	}
	// query for insert one like in the database.
	result, err := db.Collection("pebbles").InsertOne(context.Background(), pebbles)
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			handler.ResponseWriter(res, http.StatusNotAcceptable, "username or email already exists in database.", nil)
		default:
			handler.ResponseWriter(res, http.StatusInternalServerError, "Error while inserting data.", nil)
		}
		return
	}
	pebbles.ID = result.InsertedID.(primitive.ObjectID)
	handler.ResponseWriter(res, http.StatusCreated, "", pebbles)
}

// GetPebbles will handle pebbles list get request
func GetPebbles(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	//var pebblesList []schema.Pebbles

	lookupStage := bson.D{
		{
			Key: "$lookup",
			Value: bson.M{
				"from":         "creator",
				"localField":   "creator_id",
				"foreignField": "_id",
				"as":           "creator_details",
			},
		},
	}
	lookupStage2 := bson.D{
		{
			Key: "$lookup",
			Value: bson.M{
				"from":         "product",
				"localField":   "product_id",
				"foreignField": "_id",
				"as":           "product_details",
			},
		},
	}
	projectStage := bson.D{
		{
			Key: "$project",
			Value: bson.M{
				"creator_id": 0,
				"product_id": 0,
			},
		},
	}

	unwindStage := bson.D{
		{
			Key: "$unwind",
			Value: bson.M{
				"path":                       "$creator_details",
				"preserveNullAndEmptyArrays": true,
			},
		},
	}

	unwindStage2 := bson.D{
		{
			Key: "$unwind",
			Value: bson.M{
				"path":                       "$product_details",
				"preserveNullAndEmptyArrays": true,
			},
		},
	}

	pipeline := mongo.Pipeline{lookupStage, lookupStage2, projectStage, unwindStage, unwindStage2}

	showLoadedCursor, err := db.Collection("pebbles").Aggregate(context.TODO(), pipeline)
	if err != nil {
		fmt.Println("Hello", err)

	}
	var pebblesList = []bson.M{}

	if err = showLoadedCursor.All(context.TODO(), &pebblesList); err != nil {
		fmt.Println("Hellooo")

	}

	handler.ResponseWriter(res, http.StatusOK, "", pebblesList)

	// curser, err := db.Collection("pebbles").Find(context.Background(), bson.M{})
	// if err != nil {
	// 	log.Printf("Error while quering collection: %v\n", err)
	// 	handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
	// 	return
	// }
	// err = curser.All(context.Background(), &pebblesList)
	// if err != nil {
	// 	log.Fatalf("Error in curser: %v", err)
	// 	handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
	// 	return
	// }
	// handler.ResponseWriter(res, http.StatusOK, "", pebblesList)
}

func GetPebble(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		handler.ResponseWriter(res, http.StatusBadRequest, "id that you sent is wrong!!!", nil)
		return
	}
	var pebble schema.Pebbles
	// query for finding one user in the database.
	err = db.Collection("pebbles").FindOne(context.Background(), model.Pebbles{ID: id}).Decode(&pebble)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			handler.ResponseWriter(res, http.StatusNotFound, "creator not found", nil)
		default:
			log.Printf("Error while decode to go struct:%v\n", err)
			handler.ResponseWriter(res, http.StatusInternalServerError, "there is an error on server!!!", nil)
		}
		return
	}
	handler.ResponseWriter(res, http.StatusOK, "", pebble)
}
