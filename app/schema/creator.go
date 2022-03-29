package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Creator struct {
	ID          primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string                 `json:"name,omitempty" bson:"name,omitempty"`
	Email       string                 `json:"email,omitempty" bson:"email,omitempty"`
	PhoneNumber int                    `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
	Address     string                 `json:"address,omitempty" bson:"address,omitempty"`
	Category    string                 `json:"category,omitempty" bson:"category,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty" bson:"data,omitempty"` // data is a optional fields that can hold anything in key:value format.
}
