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

// CreateBrandCtaegory will handle the create brandcategory post request
func CreateBrandCategory(db *mongo.Database, res http.ResponseWriter, req *http.Request) {

	brandcategory := new(model.BrandCategory)
	err := json.NewDecoder(req.Body).Decode(brandcategory)

	result, err := db.Collection("brandcategory").InsertOne(context.TODO(), brandcategory)
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			handler.ResponseWriter(res, http.StatusNotAcceptable, "Email already exists in database.", nil)
		default:
			handler.ResponseWriter(res, http.StatusInternalServerError, "Error while inserting data.", nil)
		}
		return
	}
	brandcategory.ID = result.InsertedID.(primitive.ObjectID)
	handler.ResponseWriter(res, http.StatusCreated, "", brandcategory)
}

// GetBrandCategories will handle creator list get request
func GetBrandCategories(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var brandcategoryList []schema.BrandCategory
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
	curser, err := db.Collection("brandcategory").Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("Error while quering collection: %v\n", err)
		handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
		return
	}
	err = curser.All(context.Background(), &brandcategoryList)
	if err != nil {
		log.Fatalf("Error in curser: %v", err)
		handler.ResponseWriter(res, http.StatusInternalServerError, "Error happend while reading data", nil)
		return
	}
	handler.ResponseWriter(res, http.StatusOK, "", brandcategoryList)
}

// GetBrandcategory will give us creator with special id
func GetBrandcategory(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		handler.ResponseWriter(res, http.StatusBadRequest, "id that you sent is wrong!!!", nil)
		return
	}
	var brandcategory schema.BrandCategory
	// query for finding one user in the database.
	err = db.Collection("product").FindOne(context.Background(), model.BrandCategory{ID: id}).Decode(&brandcategory)
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
	handler.ResponseWriter(res, http.StatusOK, "", brandcategory)
}
