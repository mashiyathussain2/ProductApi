package app

import (
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
	"golang.org/x/net/context"
)

// CreateCampaignDetails will handle the create CreateCampaignDetails post request
func CreateCampaignDetail(db *mongo.Database, res http.ResponseWriter, req *http.Request) {

	campaigndetails := new(model.CampaignDetails)
	err := json.NewDecoder(req.Body).Decode(campaigndetails)

	result, err := db.Collection("campaigndetails").InsertOne(context.TODO(), campaigndetails)
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			handler.ResponseWriter(res, http.StatusNotAcceptable, "brand already exists in database.", nil)
		default:
			handler.ResponseWriter(res, http.StatusInternalServerError, "brand while inserting data.", nil)
		}
		return
	}
	campaigndetails.ID = result.InsertedID.(primitive.ObjectID)
	handler.ResponseWriter(res, http.StatusCreated, "", campaigndetails)
}

// GetCampaignDetails will handle CampaignDetail list get request
func GetCampaignDetails(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	//var campaigndetailsList []schema.CampaignDetails

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
	projectStage := bson.D{
		{
			Key: "$project",
			Value: bson.M{
				"creator_id": 0,
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

	pipeline := mongo.Pipeline{lookupStage, projectStage, unwindStage}

	showLoadedCursor, err := db.Collection("campaigndetails").Aggregate(context.TODO(), pipeline)
	if err != nil {
		fmt.Println("Hello", err)

	}
	var campaigndetailsList = []bson.M{}

	if err = showLoadedCursor.All(context.TODO(), &campaigndetailsList); err != nil {
		fmt.Println("Hellooo")

	}

	// query for find the user in the database
	// curser, err := db.Collection("campaigndetails").Find(context.Background(), bson.M{})
	// if err != nil {
	// 	log.Printf("Error while quering collection: %v\n", err)
	// 	handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", err.Error())
	// 	return
	// }
	// err = curser.All(context.Background(), &campaigndetailsList)
	// if err != nil {
	// 	log.Fatalf("Error in curser: %v", err)
	// 	handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", err.Error())
	// 	return
	// }
	handler.ResponseWriter(res, http.StatusOK, "", campaigndetailsList)
}

// GetCampaignDetail will give us CampaignDetail with special id
func GetCampaignDetail(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		handler.ResponseWriter(res, http.StatusBadRequest, "id that you sent is wrong!!!", nil)
		return
	}
	var campaigndetails schema.CampaignDetails
	// query for finding one user in the database.
	err = db.Collection("campaigndetails").FindOne(context.Background(), model.CampaignDetails{ID: id}).Decode(&campaigndetails)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			handler.ResponseWriter(res, http.StatusNotFound, "brand not found", err.Error())
		default:
			log.Printf("Error while decode to go struct:%v\n", err)
			handler.ResponseWriter(res, http.StatusInternalServerError, "there is an error on server!!!", nil)
		}
		return
	}
	handler.ResponseWriter(res, http.StatusOK, "", campaigndetails)
}
