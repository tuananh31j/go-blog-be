package common

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommonModal struct {
	Id        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	CreatedAt *time.Time         `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time         `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
