package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Date     `bson:",inline"`
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	GoogleID string             `json:"google_id,omitempty" bson:"google_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
	Picture  string             `json:"picture" bson:"picture"`
	// Likings  []string           `json:"likings" bson:"likings"`
}
