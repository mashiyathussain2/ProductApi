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

// CreateDashboard will handle the create dashboard post request
func CreateDashboard(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	dashboard := new(schema.Dashboard)
	err := json.NewDecoder(req.Body).Decode(dashboard)
	if err != nil {
		handler.ResponseWriter(res, http.StatusBadRequest, "body json request have issues!!!", nil)
		return
	}
	// query for insert one like in the database.
	result, err := db.Collection("dashboard").InsertOne(context.Background(), dashboard)
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			handler.ResponseWriter(res, http.StatusNotAcceptable, "username or email already exists in database.", nil)
		default:
			handler.ResponseWriter(res, http.StatusInternalServerError, "Error while inserting data.", nil)
		}
		return
	}
	dashboard.ID = result.InsertedID.(primitive.ObjectID)
	handler.ResponseWriter(res, http.StatusCreated, "", dashboard)
}

// GetDashboard will handle pebbles list get request
func GetDashboards(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	// var dashboardList []schema.Dashboard

	lookupStage := bson.D{
		{
			Key: "$lookup",
			Value: bson.M{
				"from":         "campaigndetails",
				"localField":   "campaign_id",
				"foreignField": "_id",
				"as":           "campaign_details",
			},
		},
	}
	lookupStage2 := bson.D{
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
	projectStage := bson.D{
		{
			Key: "$project",
			Value: bson.M{
				"campaign_id": 0,
				"creator_id":  0,
			},
		},
	}

	pipeline := mongo.Pipeline{lookupStage, lookupStage2, projectStage}

	showLoadedCursor, err := db.Collection("dashboard").Aggregate(context.TODO(), pipeline)
	if err != nil {
		fmt.Println("Hello", err)

	}
	var dashboardList = []bson.M{}

	if err = showLoadedCursor.All(context.TODO(), &dashboardList); err != nil {
		fmt.Println("Hellooo")

	}

	handler.ResponseWriter(res, http.StatusOK, "", dashboardList)

	// curser, err := db.Collection("dashboard").Find(context.Background(), bson.M{})
	// if err != nil {
	// 	log.Printf("Error while quering collection: %v\n", err)
	// 	handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
	// 	return
	// }
	// err = curser.All(context.Background(), &dashboardList)
	// if err != nil {
	// 	log.Fatalf("Error in curser: %v", err)
	// 	handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
	// 	return
	// }
	// handler.ResponseWriter(res, http.StatusOK, "", dashboardList)
}

func GetDashboard(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		handler.ResponseWriter(res, http.StatusBadRequest, "id that you sent is wrong!!!", nil)
		return
	}
	var dashboard schema.Dashboard
	// query for finding one user in the database.
	err = db.Collection("dashboard").FindOne(context.Background(), model.Dashboard{ID: id}).Decode(&dashboard)
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
	handler.ResponseWriter(res, http.StatusOK, "", dashboard)
}
