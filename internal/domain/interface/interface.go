package repository

import (
	"context"
	"mime/multipart"

	blogModel "nta-blog/internal/domain/model/blog"
	guestbookModel "nta-blog/internal/domain/model/guestBook"
	imageModel "nta-blog/internal/domain/model/image"
	tagModel "nta-blog/internal/domain/model/tag"
	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogRepository interface {
	Create(ctx context.Context, dto *blogModel.Blog) error
	Find(ctx context.Context, conditions map[string]interface{}) (*blogModel.Blog, error)
	List(ctx context.Context, pipeline bson.A) ([]blogModel.Blog, error)
	Update(ctx context.Context, conditions map[string]interface{}, dto map[string]interface{}) error
}

type UserRepository interface {
	CreateUser(ctx context.Context, dto *userModel.User) error
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
	List(ctx context.Context, pipeline bson.A) ([]userModel.User, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) error
}

type GuestRepository interface {
	Create(ctx context.Context, dto *guestbookModel.GuestBook) error
	Find(ctx context.Context, condition map[string]interface{}) ([]*guestbookModel.GuestBook, error)
	List(ctx context.Context, pipeline bson.A) ([]guestbookModel.GuestBook, error)
	Update(ctx context.Context, guestBookId primitive.ObjectID, updateField map[string]interface{}) error
}

type ImageRepository interface {
	ListImages(ctx context.Context) ([]imageModel.Image, error)
	SaveImage(ctx context.Context, image *imageModel.Image) error
	UploadFile(ctx context.Context, file multipart.File) (*imageModel.UploadResFormCld, error)
}

type TagRepository interface {
	Create(ctx context.Context, dto *tagModel.Tag) error
	Find(ctx context.Context, conditions map[string]interface{}) (*tagModel.Tag, error)
	List(ctx context.Context, conditions map[string]interface{}) ([]tagModel.Tag, error)
	Update(ctx context.Context, condition map[string]interface{}, dto map[string]interface{}) error
}

type TokenRepository interface {
	FindToken(ctx context.Context, userId string) (string, error)
	RemoveToken(ctx context.Context, userId string) error
	SaveRefreshToken(ctx context.Context, token, userId string) error
}
