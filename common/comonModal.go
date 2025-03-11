package common

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommonModal struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt *time.Time         `bson:"created_at" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
}
