package blogService

import (
	"context"

	blogModel "nta-blog/internal/domain/model/blog"

	"go.mongodb.org/mongo-driver/bson"
)

type DetailsBlogStore interface {
	List(ctx context.Context, pipeline bson.A) ([]blogModel.Blog, error)
	Find(ctx context.Context, conditions map[string]interface{}) (*blogModel.Blog, error)
}

type detailsBlogService struct {
	blogStore DetailsBlogStore
}

func NewDetailsBlogService(blogStore DetailsBlogStore) *detailsBlogService {
	return &detailsBlogService{blogStore: blogStore}
}

func (sv *detailsBlogService) FindDetailsBlog(ctx context.Context, conditions map[string]interface{}) (*blogModel.Blog, error) {
	return sv.blogStore.Find(ctx, conditions)
}

func (sv *detailsBlogService) ListBlog(ctx context.Context, pipeline bson.A) ([]blogModel.Blog, error) {
	return sv.blogStore.List(ctx, pipeline)
}
