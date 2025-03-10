package db

import (
	"context"
	"log"
	"sync"
	"time"

	"nta-blog/common"
	cnst "nta-blog/constant"
	"nta-blog/libs/hashser"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoOnce      sync.Once
	mongoClient    *mongo.Client
	mongoClientErr error
)

func ConnectMongo(uri string) (*mongo.Client, error) {
	mongoOnce.Do(func() {
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(uri).SetMaxPoolSize(100).SetConnectTimeout(100 * time.Second).SetServerAPIOptions(serverAPI)
		ctx := context.Background()
		client, err := mongo.Connect(ctx, opts)
		if err != nil {
			mongoClientErr = err
		} else {
			err = client.Ping(ctx, nil)
			if err != nil {
				mongoClientErr = err
			}
		}
		mongoClient = client
	})

	return mongoClient, mongoClientErr
}

func SetupUserCollection(db *mongo.Database) {
	userDB := db.Collection(cnst.UserCollection)
	salt := common.GenSalt()
	password := "12345678"
	md5Hash := hashser.NewMd5Hash(password)
	md5Hash.SetSalt(salt)
	hash := md5Hash.Hash()
	if _, err := userDB.DeleteMany(context.Background(), bson.D{}); err != nil {
		log.Fatal(err)
	}
	createIndexFiled(userDB, "email")

	_, err := userDB.InsertOne(context.Background(), bson.D{
		{Key: "email", Value: "admin@gmail.com"},
		{Key: "role", Value: 1},
		{Key: "password", Value: hash},
		{Key: "salt", Value: salt},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func SetupBlogCollection(db *mongo.Database) {
	blogDB := db.Collection(cnst.BlogCollection)
	createIndexFiled(blogDB, "tags")
}

func SetupTagCollection(db *mongo.Database) {
	db.Collection(cnst.TagCollection)
}

func createIndexFiled(col *mongo.Collection, field string) {
	_, err := col.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: field, Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		panic(err)
	}
}

func DisconnectMongo(i *mongo.Client) {
	if err := i.Disconnect(context.Background()); err != nil {
		log.Fatalf("Somethings wrong when disconnect mongodb: >>>>>>>>>%v", err)
	}
}
