package blogBusiness

import (
	"context"

	blogModel "nta-blog/internal/domain/model/blog"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateBlogService interface {
	CheckTagExists(ctx context.Context, tapId primitive.ObjectID) error
	CheckUserExists(ctx context.Context, userId primitive.ObjectID) error
	Create(ctx context.Context, dto *blogModel.Blog) error
}

type createBlogBiz struct {
	service CreateBlogService
}

func NewCreateBlogBiz(service CreateBlogService) *createBlogBiz {
	return &createBlogBiz{service: service}
}

func (biz *createBlogBiz) CreateBlog(ctx context.Context, dto *blogModel.Blog) error {
	for _, tagId := range dto.TagIds {
		if err := biz.service.CheckTagExists(ctx, tagId); err != nil {
			return err
		}
	}
	if err := biz.service.CheckUserExists(ctx, dto.UserId); err != nil {
		return err
	}
	return biz.service.Create(ctx, dto)
}
