package db

import (
	"context"
	"log"
	"sync"
	"time"

	blogModel "nta-blog/internal/domain/model/blog"
	guestbookModel "nta-blog/internal/domain/model/guestBook"
	imageModel "nta-blog/internal/domain/model/image"
	tagModel "nta-blog/internal/domain/model/tag"
	userModel "nta-blog/internal/domain/model/user"

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
	userDB := db.Collection(userModel.UserCollectionName)

	createIndexFiled(userDB, "email")
}

func SetupBlogCollection(db *mongo.Database) {
	blogDB := db.Collection(blogModel.BlogCollection)
	createIndexFiledNotUnique(blogDB, "tag_ids")
}

func SetupImageCollection(db *mongo.Database) {
	db.Collection(imageModel.ImageCollection)
}

func SetupGuestBookCollection(db *mongo.Database) {
	guestBookDB := db.Collection(guestbookModel.GuestBookCollection)
	createIndexFiledNotUnique(guestBookDB, "user_id")
	// if _, err := guestBookDB.DeleteMany(context.Background(), bson.D{}); err != nil {
	// 	log.Fatal(err)
	// }
}

func SetupTagCollection(db *mongo.Database) {
	db.Collection(tagModel.TagCollection)
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

func createIndexFiledNotUnique(col *mongo.Collection, field string) {
	_, err := col.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: field, Value: 1}},
		Options: options.Index().SetUnique(false),
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
