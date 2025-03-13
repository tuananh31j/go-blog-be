package blogService

import (
	"context"

	blogModel "nta-blog/internal/domain/model/blog"
	tagModel "nta-blog/internal/domain/model/tag"
	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateBlogTagStore interface {
	Find(ctx context.Context, conditions map[string]interface{}) (*tagModel.Tag, error)
}

type CreateBlogUserStore interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
}

type CreateBlogService interface {
	Create(ctx context.Context, dto *blogModel.Blog) error
}

type createBlogService struct {
	tagStore  CreateBlogTagStore
	userStore CreateBlogUserStore
	blogStore CreateBlogService
}

func NewCreateBlogService(tagStore CreateBlogTagStore, userStore CreateBlogUserStore, blogStore CreateBlogService) *createBlogService {
	return &createBlogService{tagStore: tagStore, userStore: userStore, blogStore: blogStore}
}

func (sv *createBlogService) CheckTagExists(ctx context.Context, tapId primitive.ObjectID) error {
	_, err := sv.tagStore.Find(ctx, map[string]interface{}{"_id": tapId})
	return err
}

func (sv *createBlogService) CheckUserExists(ctx context.Context, userId primitive.ObjectID) error {
	_, err := sv.userStore.FindOneUser(ctx, map[string]interface{}{"_id": userId})
	return err
}

func (sv *createBlogService) Create(ctx context.Context, dto *blogModel.Blog) error {
	return sv.blogStore.Create(ctx, dto)
}
