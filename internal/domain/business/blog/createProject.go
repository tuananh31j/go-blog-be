package blogBusiness

import (
	"context"
	"time"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	blogModel "nta-blog/internal/domain/model/blog"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateProjectService interface {
	CheckTagExists(ctx context.Context, tapId primitive.ObjectID) error
	CheckUserExists(ctx context.Context, userId primitive.ObjectID) error
	Create(ctx context.Context, dto *blogModel.Blog) error
}

type createProjectBiz struct {
	service CreateProjectService
}

func NewCreateProjectBiz(service CreateProjectService) *createProjectBiz {
	return &createProjectBiz{service: service}
}

func (biz *createProjectBiz) CreateProject(ctx context.Context, dto *blogModel.CreatePayload) error {
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
	blog.Type = cnst.BlogTypeConstant.Project

	if err := biz.service.CheckUserExists(ctx, dto.UserId); err != nil {
		return err
	}
	err := biz.service.Create(ctx, &blog)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
