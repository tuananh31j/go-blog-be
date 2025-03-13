package blogService

import (
	"context"

	blogModel "nta-blog/internal/domain/model/blog"

	"go.mongodb.org/mongo-driver/bson"
)

type ListBlogStore interface {
	List(ctx context.Context, pipeline bson.A) ([]blogModel.Blog, error)
}

type listBlogService struct {
	blogStore ListBlogStore
}

func NewListBlogStore(blogStore ListBlogStore) *listBlogService {
	return &listBlogService{blogStore: blogStore}
}

func (sv *listBlogService) List(ctx context.Context, pipeline bson.A) ([]blogModel.Blog, error) {
	return sv.blogStore.List(ctx, pipeline)
}
