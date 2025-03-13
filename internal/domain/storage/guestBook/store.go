package guestBookStorage

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type store struct {
	db *mongo.Database
}

func NewStore(db *mongo.Database) *store {
	return &store{db: db}
}
