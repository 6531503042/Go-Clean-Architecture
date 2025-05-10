package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User struct {
		ID          primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
		Email       string              `json:"email" bson:"email,omitempty"`
		Password    string              `json:"password" bson:"password,omitempty"`
		FirstName   string              `json:"first_name" bson:"first_name,omitempty"`
		LastName    string              `json:"last_name" bson:"last_name,omitempty"`
		PhoneNumber string              `json:"phone_number" bson:"phone_number,omitempty"`
		CreatedAt   time.Time           `json:"created_at" bson:"created_at,omitempty"`
		UpdatedAt   time.Time           `json:"updated_at" bson:"updated_at,omitempty"`
	}
)