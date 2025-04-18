package book

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct{
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `json:"title"`
    Author    string             `json:"author"`
    Genre     string             `json:"genre"`
    Price     float64            `json:"price"`
    InStock   bool               `json:"in_stock"`
    CreatedAt time.Time          `json:"created_at"`
}