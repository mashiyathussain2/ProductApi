package app

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"

	"productlist/app/handler"
	"productlist/app/model"
	"productlist/app/schema"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

// CreateBrand will handle the create brand post request
func CreateBrand(db *mongo.Database, res http.ResponseWriter, req *http.Request) {

	brand := new(model.Brand)
	err := json.NewDecoder(req.Body).Decode(brand)

	result, err := db.Collection("brand").InsertOne(context.TODO(), brand)
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			handler.ResponseWriter(res, http.StatusNotAcceptable, "brand already exists in database.", nil)
		default:
			handler.ResponseWriter(res, http.StatusInternalServerError, "brand while inserting data.", nil)
		}
		return
	}
	brand.ID = result.InsertedID.(primitive.ObjectID)
	handler.ResponseWriter(res, http.StatusCreated, "", brand)
}

// GetBrands will handle product list get request
func GetBrands(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var brandList []schema.Brand
	pageString := req.FormValue("page")
	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		page = 0
	}
	page = page * limit
	findOptions := options.FindOptions{
		Skip:  &page,
		Limit: &limit,
		Sort: bson.M{
			"_id": -1, // -1 for descending and 1 for ascending
		},
	}
	// query for find the user in the database
	curser, err := db.Collection("brand").Find(context.Background(), bson.M{}, &findOptions)
	if err != nil {
		log.Printf("Error while quering collection: %v\n", err)
		handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
		return
	}
	err = curser.All(context.Background(), &brandList)
	if err != nil {
		log.Fatalf("Error in curser: %v", err)
		handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
		return
	}
	handler.ResponseWriter(res, http.StatusOK, "", brandList)
}

// GetBrand will give us brand with special id
func GetBrand(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		handler.ResponseWriter(res, http.StatusBadRequest, "id that you sent is wrong!!!", nil)
		return
	}
	var brand schema.Brand
	// query for finding one user in the database.
	err = db.Collection("brand").FindOne(context.Background(), model.Brand{ID: id}).Decode(&brand)
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
	handler.ResponseWriter(res, http.StatusOK, "", brand)
}
