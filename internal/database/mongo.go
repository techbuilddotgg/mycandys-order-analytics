package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Db *mongo.Database

func Connect(uri string) *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	Db = client.Database("analytics")
	return client
}

func Disconnect(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
