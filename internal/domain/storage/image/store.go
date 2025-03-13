package imageStorage

import (
	"github.com/cloudinary/cloudinary-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type store struct {
	db  *mongo.Database
	cld *cloudinary.Cloudinary
}

func NewStore(db *mongo.Database, cld *cloudinary.Cloudinary) *store {
	return &store{db: db, cld: cld}
}
