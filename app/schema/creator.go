package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Creator struct {
	ID              primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            string                 `json:"name,omitempty" bson:"name,omitempty"`
	CampaignName    string                 `json:"campaign_name,omitempty" bson:"campaign_name,omitempty"`
	Email           string                 `json:"email,omitempty" bson:"email,omitempty"`
	PhoneNumber     int                    `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
	Address         string                 `json:"address,omitempty" bson:"address,omitempty"`
	Category        string                 `json:"category,omitempty" bson:"category,omitempty"`
	ActivationStore bool                   `json:"activation_store,omitempty" bson:"activation_store,omitempty"`
	StoreLink       string                 `json:"store_link,omitempty" bson:"store_link,omitempty"`
	BudgetProvided  int                    `json:"budget_provided,omitempty" bson:"budget_provided,omitempty"`
	Gmv             int                    `json:"gmv,omitempty" bson:"gmv,omitempty"`
	Data            map[string]interface{} `json:"data,omitempty" bson:"data,omitempty"` // data is a optional fields that can hold anything in key:value format.
}
