package appctx

import (
	"github.com/cloudinary/cloudinary-go"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	GetMongoDB() *mongo.Database
	GetRedis() *redis.Client
	GetLogger() *zerolog.Logger
	GetCloudinary() *cloudinary.Cloudinary
}

type appctx struct {
	db     *mongo.Database
	rdb    *redis.Client
	cld    *cloudinary.Cloudinary
	logger *zerolog.Logger
}

func NewAppContext(db *mongo.Database, rdb *redis.Client, cld *cloudinary.Cloudinary, lg *zerolog.Logger) *appctx {
	return &appctx{db: db, logger: lg, cld: cld, rdb: rdb}
}

func (actx *appctx) GetMongoDB() *mongo.Database {
	return actx.db
}

func (actx *appctx) GetRedis() *redis.Client {
	return actx.rdb
}

func (actx *appctx) GetCloudinary() *cloudinary.Cloudinary {
	return actx.cld
}

func (actx *appctx) GetLogger() *zerolog.Logger {
	return actx.logger
}
