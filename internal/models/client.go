package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/websocket"
)

type Client struct {
	ID   primitive.ObjectID `bson:"id"`
	Conn *websocket.Conn
	Hub  *Hub
	User User
	send chan []byte
}
