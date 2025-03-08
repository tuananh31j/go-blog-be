package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(uri string) *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetMaxPoolSize(100).SetConnectTimeout(100 * time.Second).SetServerAPIOptions(serverAPI)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err.Error())
	}
	return client
}

func DisconnectMongo(i *mongo.Client) {
	if err := i.Disconnect(context.Background()); err != nil {
		log.Fatalf("Somethings wrong when disconnect mongodb: >>>>>>>>>%v", err)
	}
}
