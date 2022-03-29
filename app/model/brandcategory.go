package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BrandCategory struct {
	ID        primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Category  string                 `json:"category,omitempty" bson:"category,omitempty"`
	BrandName []string               `json:"brand_name,omitempty" bson:"brand_name,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty" bson:"data,omitempty"` // data is a optional fields that can hold anything in key:value format.
}
