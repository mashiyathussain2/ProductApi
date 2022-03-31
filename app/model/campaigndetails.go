package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CampaignDetails struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Type      string             `json:"type,omitempty" bson:"type,omitempty"`
	StartDate string             `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate   string             `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Products  []string           `json:"products,omitempty" bson:"products,omitempty"`
	//BrandsChosen   []string               `json:"brands_chosen,omitempty" bson:"brands_chosen,omitempty"` // remove this in future
	CreatorDetails primitive.ObjectID     `json:"creator_id,omitempty" bson:"creator_id,omitempty"`
	Data           map[string]interface{} `json:"data,omitempty" bson:"data,omitempty"` // data is a optional fields that can hold anything in key:value format.
}
