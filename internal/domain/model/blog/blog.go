package blogModel

import (
	"time"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	tagModel "nta-blog/internal/domain/model/tag"
	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const BlogCollection = "blogs"

type Blog struct {
	common.CommonModal `bson:",inline"`
	Thumbnail          string                 `bson:"thumbnail" json:"thumbnail"`
	Summary            string                 `bson:"summary" json:"summary"`
	Title              string                 `bson:"title" json:"title"`
	Content            string                 `bson:"content" json:"content"`
	UserId             primitive.ObjectID     `bson:"user_id" json:"user_id"`
	User               *userModel.PrivateUser `bson:",omitempty" json:"user,omitempty"`
	Status             cnst.TStatusBlog       `bson:"status" json:"status"`
	TagIds             []primitive.ObjectID   `bson:"tag_ids"`
	Type               cnst.IBlogType         `bson:"type"`
	Tags               []tagModel.Tag         `bson:",omitempty" json:"tags,omitempty"`
}

type CreatePayload struct {
	Thumbnail string `json:"thumbnail"`
	Summary   string `json:"summary"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserId    primitive.ObjectID
	TagIds    []string         `json:"tag_ids"`
	Status    cnst.TStatusBlog `json:"status"`
}
type UpdatePayload struct {
	Thumbnail string           `json:"thumbnail"`
	Summary   string           `json:"summary"`
	Title     string           `json:"title"`
	Content   string           `json:"content"`
	TagIds    []string         `json:"tag_ids"`
	Status    cnst.TStatusBlog `json:"status"`
	Type      cnst.IBlogType   `json:"type"`
	UpdatedAt *time.Time       `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
