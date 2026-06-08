package models

import (
	"github.com/carlosEA28/go_chat/internal/repositories"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hub struct {
	ID           primitive.ObjectID `bson:"id"`
	Clients      map[*Client]bool
	Register     chan *Client
	Unregister   chan *Client
	Brodcast     chan Message
	Subcriptions map[string]bool
	mongo        *repositories.MongoRepositoryContext
	redis        *redis.Client
}
