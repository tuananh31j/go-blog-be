package blogBusiness

import (
	"context"
	"time"

	"nta-blog/internal/common"
	blogModel "nta-blog/internal/domain/model/blog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateService interface {
	CheckTagExists(ctx context.Context, tapId primitive.ObjectID) error
	FindDetailsBlog(ctx context.Context, conditions map[string]interface{}) (*blogModel.Blog, error)
	Update(ctx context.Context, conditions map[string]interface{}, dto map[string]interface{}) error
}

type updateBlogBiz struct {
	service UpdateService
}

func NewUpdateBiz(service UpdateService) *updateBlogBiz {
	return &updateBlogBiz{service: service}
}

func (biz *updateBlogBiz) Update(ctx context.Context, blogId primitive.ObjectID, dto *blogModel.UpdatePayload) error {
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

	// blog.TagIds = objTagIsd
	// blog.Thumbnail = dto.Thumbnail
	// blog.Summary = dto.Summary
	// blog.Title = dto.Title
	// blog.Content = dto.Content
	// blog.Status = dto.Status
	// blog.UpdatedAt = &now
	// blog.Type = dto.Type
	updateData := bson.M{"$set": bson.M{
		"title":      dto.Title,
		"content":    dto.Content,
		"thumnail":   dto.Thumbnail,
		"summary":    dto.Summary,
		"status":     dto.Status,
		"tag_ids":    dto.TagIds,
		"type":       dto.Type,
		"updated_at": &now,
	}}

	_, err := biz.service.FindDetailsBlog(ctx, map[string]interface{}{"_id": blogId})
	if err != nil {
		return common.ErrBadRequest(err)
	}

	if err := biz.service.Update(ctx, map[string]interface{}{"_id": blogId}, updateData); err != nil {
		return common.ErrInternal(err)
	}

	return nil
}
