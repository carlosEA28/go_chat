package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID        primitive.ObjectID `bson:"id"`
	Username  string             `bson:"username"`
	Deparment string             `bson:"department"`
	Content   string             `bson:"content"`
	CreatedAt primitive.DateTime `bson:"created_at"`
}
