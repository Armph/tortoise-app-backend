package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// pet model from pet collection
type Pet struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name           string             `json:"name"`
	Age            int32              `json:"age"`
	Price          int32              `json:"price"`
	Is_sold        bool               `json:"is_sold"`
	Description    string             `json:"description"`
	Weight         int32              `json:"weight"`
	Sex            string             `json:"sex"`
	Species        string             `json:"species"`
	Type           string             `json:"type"`
	Behavior       string             `json:"behavior"`
	Media          string             `json:"media"`
	Medical_record struct {
		Medical_id  string    `json:"medical_id"`
		Date        time.Time `json:"date"`
		Description string    `json:"description"`
	} `json:"medical_record"`
	Seller_id primitive.ObjectID `bson:"seller_id,omitempty" json:"seller_id,omitempty"`
}

// pet card model from pet collection
type PetCard struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name           string             `json:"name"`
	Price          int32              `json:"price"`
	Type           string             `json:"type"`
	Media          string             `json:"media"`
	Seller_id      string             `json:"seller_id"`
	Seller_name    string             `json:"seller_name"`
	Seller_surname string             `json:"seller_surname"`
}