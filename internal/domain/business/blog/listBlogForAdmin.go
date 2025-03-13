package blogBusiness

import (
	"context"

	blogModel "nta-blog/internal/domain/model/blog"

	"go.mongodb.org/mongo-driver/bson"
)

type ListBlogService interface {
	List(ctx context.Context, pipeline bson.A) ([]blogModel.Blog, error)
}

type listBlogBiz struct {
	service ListBlogService
}

func NewListBlogBiz(service ListBlogService) *listBlogBiz {
	return &listBlogBiz{service: service}
}

func (biz *listBlogBiz) ListBlog(ctx context.Context, pipeline bson.A) ([]blogModel.Blog, error) {
	return biz.service.List(ctx, pipeline)
}
