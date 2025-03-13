package blogService

import (
	"context"

	blogModel "nta-blog/internal/domain/model/blog"
)

type DetailsBlogStore interface {
	Find(ctx context.Context, conditions map[string]interface{}) (*blogModel.Blog, error)
}

type detailsBlogService struct {
	blogStore DetailsBlogStore
}

func NewDetailsBlogStore(blogStore DetailsBlogStore) *detailsBlogService {
	return &detailsBlogService{blogStore: blogStore}
}

func (sv *detailsBlogService) FindDetailsBlog(ctx context.Context, conditions map[string]interface{}) (*blogModel.Blog, error) {
	return sv.blogStore.Find(ctx, conditions)
}
