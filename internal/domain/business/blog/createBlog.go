package blogBusiness

import (
	"context"
	"time"

	"nta-blog/internal/common"
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

func (biz *createBlogBiz) CreateBlog(ctx context.Context, dto *blogModel.CreateBlogPayload) error {
	var blog blogModel.Blog
	now := time.Now()

	objTagIsd := make([]primitive.ObjectID, 0)

	for _, tagId := range dto.TagIds {
		objId, err := primitive.ObjectIDFromHex(tagId)
		if err != nil {
			return common.ErrInvalidRequest(err)
		}
		// if err := biz.service.CheckTagExists(ctx, objId); err != nil {
		// 	return err
		// }
		objTagIsd = append(objTagIsd, objId)
	}

	blog.TagIds = objTagIsd
	blog.Thumbnail = dto.Thumbnail
	blog.Summary = dto.Summary
	blog.Title = dto.Title
	blog.Content = dto.Content
	blog.UserId = dto.UserId
	blog.Status = dto.Status
	blog.CreatedAt = &now
	blog.UpdatedAt = &now
	blog.Id = primitive.NewObjectID()

	if err := biz.service.CheckUserExists(ctx, dto.UserId); err != nil {
		return err
	}
	err := biz.service.Create(ctx, &blog)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
