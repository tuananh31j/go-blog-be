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
	Thumbnail          string               `bson:"thumbnail" json:"thumbnail"`
	Summary            string               `bson:"summary" json:"summary"`
	Title              string               `bson:"title" json:"title"`
	Content            string               `bson:"content" json:"content"`
	UserId             primitive.ObjectID   `bson:"user_id" json:"user_id"`
	Status             cnst.TStatusBlog     `bson:"status" json:"status"`
	TagIds             []primitive.ObjectID `bson:"tag_ids"`
	Tags               []tagModel.Tag       `json:"tags"`
}

type CreateBlogPayload struct {
	Thumbnail string `json:"thumbnail"`
	Summary   string `json:"summary"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserId    primitive.ObjectID
	TagIds    []string         `json:"tag_ids"`
	Status    cnst.TStatusBlog `json:"status"`
}
