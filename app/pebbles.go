package app

import (
	"context"
	"encoding/json"
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
	var pebblesList []schema.Pebbles

	curser, err := db.Collection("pebbles").Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("Error while quering collection: %v\n", err)
		handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
		return
	}
	err = curser.All(context.Background(), &pebblesList)
	if err != nil {
		log.Fatalf("Error in curser: %v", err)
		handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
		return
	}
	handler.ResponseWriter(res, http.StatusOK, "", pebblesList)
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
