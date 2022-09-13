package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookModel struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Author string             `bson:"author,omitempty"`
	Pbdate string             `bson:"pbdate,omitempty"`
}
