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

var limit int64 = 10

// CreateProduct will handle the create product post request
func CreateProduct(db *mongo.Database, res http.ResponseWriter, req *http.Request) {

	product := new(model.Product)
	err := json.NewDecoder(req.Body).Decode(product)

	// First find the user with their email in database if the user already created then return already exists.
	//err = db.Collection("product").FindOne(context.Background(), model.Product{BrandName: product.BrandName}).Decode(&product)
	// if user not exists in the database then create a new user and insert that user in the database.
	result, err := db.Collection("product").InsertOne(context.TODO(), product)
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			handler.ResponseWriter(res, http.StatusNotAcceptable, "Email already exists in database.", nil)
		default:
			handler.ResponseWriter(res, http.StatusInternalServerError, "Error while inserting data.", nil)
		}
		return
	}
	product.ID = result.InsertedID.(primitive.ObjectID)
	handler.ResponseWriter(res, http.StatusCreated, "", product)
}

// GetProducts will handle product list get request
func GetProducts(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var productList []schema.Product
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
	curser, err := db.Collection("product").Find(nil, bson.M{}, &findOptions)
	if err != nil {
		log.Printf("Error while quering collection: %v\n", err)
		handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
		return
	}
	err = curser.All(context.Background(), &productList)
	if err != nil {
		log.Fatalf("Error in curser: %v", err)
		handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
		return
	}
	handler.ResponseWriter(res, http.StatusOK, "", productList)
}

// GetProduct will give us product with special id
func GetProduct(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		handler.ResponseWriter(res, http.StatusBadRequest, "id that you sent is wrong!!!", nil)
		return
	}
	var product schema.Product
	// query for finding one user in the database.
	err = db.Collection("product").FindOne(context.Background(), model.Product{ID: id}).Decode(&product)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			handler.ResponseWriter(res, http.StatusNotFound, "person not found", nil)
		default:
			log.Printf("Error while decode to go struct:%v\n", err)
			handler.ResponseWriter(res, http.StatusInternalServerError, "there is an error on server!!!", nil)
		}
		return
	}
	handler.ResponseWriter(res, http.StatusOK, "", product)
}
