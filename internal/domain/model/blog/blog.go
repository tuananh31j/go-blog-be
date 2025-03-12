package blogModel

import (
	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	tagModel "nta-blog/internal/domain/model/tag"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const BlogCollection = "blogs"

type Blog struct {
	common.CommonModal `bson:",inline"`
	Title              string               `bson:"title" json:"title"`
	Content            string               `bson:"content" json:"content"`
	UserId             primitive.ObjectID   `bson:"user_id" json:"user_id"`
	Status             cnst.Status          `bson:"status" json:"status"`
	TagIds             []primitive.ObjectID `bson:"tag_ids"`
	Tags               []tagModel.Tag       `json:"tags"`
}
