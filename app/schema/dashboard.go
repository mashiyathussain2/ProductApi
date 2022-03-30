package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dashboard struct {
	ID             primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	ManagerName    string                 `json:"manager_name,omitempty" bson:"manager_name,omitempty"`
	CampaignDetail primitive.ObjectID     `json:"campaign_id,omitempty" bson:"campaign_id,omitempty"`
	CreatorDetails primitive.ObjectID     `json:"creator_id,omitempty" bson:"creator_id,omitempty"`
	Deliverables   []string               `json:"deliverables,omitempty" bson:"deliverables,omitempty"`
	Data           map[string]interface{} `json:"data,omitempty" bson:"data,omitempty"` // data is a optional fields that can hold anything in key:value format.
}
