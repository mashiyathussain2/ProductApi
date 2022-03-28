package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductName string                 `json:"product_name,omitempty" bson:"product_name,omitempty"`
	Price       int                    `json:"price,omitempty" bson:"price,omitempty"`
	BrandName   string                 `json:"brand_name,omitempty" bson:"brand_name,omitempty"`
	Category    string                 `json:"category,omitempty" bson:"category,omitempty"`
	ProductImg  string                 `json:"product_img,omitempty" bson:"product_img,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty" bson:"data,omitempty"` // data is a optional fields that can hold anything in key:value format.
}
