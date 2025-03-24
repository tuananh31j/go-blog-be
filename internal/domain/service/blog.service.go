package service

import (
	"context"

	repository "nta-blog/internal/domain/interface"
	blogModel "nta-blog/internal/domain/model/blog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type blogService struct {
	tagStore  repository.TagRepository
	userStore repository.UserRepository
	blogStore repository.BlogRepository
}

func NewBlogService(ts repository.TagRepository, us repository.UserRepository, bs repository.BlogRepository) *blogService {
	return &blogService{tagStore: ts, userStore: us, blogStore: bs}
}

func (sv *blogService) CheckTagExists(ctx context.Context, tapId primitive.ObjectID) error {
	_, err := sv.tagStore.Find(ctx, map[string]interface{}{"_id": tapId})
	return err
}

func (sv *blogService) CheckUserExists(ctx context.Context, userId primitive.ObjectID) error {
	_, err := sv.userStore.FindOneUser(ctx, map[string]interface{}{"_id": userId})
	return err
}

func (sv *blogService) Create(ctx context.Context, dto *blogModel.Blog) error {
	return sv.blogStore.Create(ctx, dto)
}

func (sv *blogService) FindDetailsBlog(ctx context.Context, conditions map[string]interface{}) (*blogModel.Blog, error) {
	return sv.blogStore.Find(ctx, conditions)
}

func (sv *blogService) ListBlog(ctx context.Context, pipeline bson.A) ([]blogModel.Blog, error) {
	return sv.blogStore.List(ctx, pipeline)
}

func (sv *blogService) List(ctx context.Context, pipeline bson.A) ([]blogModel.Blog, error) {
	return sv.blogStore.List(ctx, pipeline)
}
