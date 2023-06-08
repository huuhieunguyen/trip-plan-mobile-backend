package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Date struct {
	// DateSort   primitive.DateTime `bson:"dateSort" json:"DateSort"`
	DateCreated primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	// DateUpdated primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	DateUpdated primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
}
