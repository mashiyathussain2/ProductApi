package app

import (
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
	"golang.org/x/net/context"
)

// CreateCreator will handle the create creator post request
func CreateCreator(db *mongo.Database, res http.ResponseWriter, req *http.Request) {

	creator := new(model.Creator)
	err := json.NewDecoder(req.Body).Decode(creator)

	// First find the user with their email in database if the user already created then return already exists.
	//err = db.Collection("product").FindOne(context.Background(), model.Product{BrandName: product.BrandName}).Decode(&product)
	// if user not exists in the database then create a new user and insert that user in the database.
	result, err := db.Collection("creator").InsertOne(context.TODO(), creator)
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			handler.ResponseWriter(res, http.StatusNotAcceptable, "Email already exists in database.", nil)
		default:
			handler.ResponseWriter(res, http.StatusInternalServerError, "Error while inserting data.", nil)
		}
		return
	}
	creator.ID = result.InsertedID.(primitive.ObjectID)
	handler.ResponseWriter(res, http.StatusCreated, "", creator)
}

// GetCreators will handle creator list get request
func GetCreators(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var creatorList []schema.Creator
	// pageString := req.FormValue("page")
	// page, err := strconv.ParseInt(pageString, 10, 64)
	// if err != nil {
	// 	page = 0
	// }
	// page = page * limit
	// findOptions := options.FindOptions{
	// 	Skip:  &page,
	// 	Limit: &limit,
	// 	Sort: bson.M{
	// 		"_id": -1, // -1 for descending and 1 for ascending
	// 	},
	// }
	// query for find the user in the database
	curser, err := db.Collection("creator").Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("Error while quering collection: %v\n", err)
		handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
		return
	}
	err = curser.All(context.Background(), &creatorList)
	if err != nil {
		log.Fatalf("Error in curser: %v", err)
		handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
		return
	}
	handler.ResponseWriter(res, http.StatusOK, "", creatorList)
}

// GetCreator will give us creator with special id
func GetCreator(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		handler.ResponseWriter(res, http.StatusBadRequest, "id that you sent is wrong!!!", nil)
		return
	}
	var creator schema.Creator
	// query for finding one user in the database.
	err = db.Collection("product").FindOne(context.Background(), model.Creator{ID: id}).Decode(&creator)
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
	handler.ResponseWriter(res, http.StatusOK, "", creator)
}
