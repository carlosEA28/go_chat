package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoRepositoryContext struct {
	*mongo.Collection
	client *mongo.Client
}

func NewMongoRepositoryContext(uri, dbName, collectionName string) (*MongoRepositoryContext, error) {

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &MongoRepositoryContext{
		Collection: collection,
		client:     client,
	}, nil
}

func (r *MongoRepositoryContext) Create(ctx context.Context, document interface{}) error {
	_, err := r.Collection.InsertOne(ctx, document)
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoRepositoryContext) Close(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}
