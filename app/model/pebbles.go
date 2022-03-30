package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pebbles struct {
	ID         primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Creator_ID primitive.ObjectID     `json:"creator_id,omitempty" bson:"creator_id,omitempty"`
	Product_ID primitive.ObjectID     `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Data       map[string]interface{} `json:"data,omitempty" bson:"data,omitempty"` // data is a optional fields that can hold anything in key:value format.
}
