package tokenStorage

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type store struct {
	db  *mongo.Database
	rdb *redis.Client
}

func NewStore(db *mongo.Database, rdb *redis.Client) *store {
	return &store{db: db, rdb: rdb}
}
