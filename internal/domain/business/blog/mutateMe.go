package blogBusiness

import (
	"context"
	"time"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	blogModel "nta-blog/internal/domain/model/blog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MutateMeService interface {
	CheckTagExists(ctx context.Context, tapId primitive.ObjectID) error
	CheckUserExists(ctx context.Context, userId primitive.ObjectID) error
	Create(ctx context.Context, dto *blogModel.Blog) error
	FindDetailsBlog(ctx context.Context, conditions map[string]interface{}) (*blogModel.Blog, error)
	Update(ctx context.Context, conditions map[string]interface{}, dto map[string]interface{}) error
}

type mutateMeBiz struct {
	service MutateMeService
}

func NewMutateMeBiz(service MutateMeService) *mutateMeBiz {
	return &mutateMeBiz{service: service}
}

func (biz *mutateMeBiz) MutateMe(ctx context.Context, dto *blogModel.CreatePayload) error {
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
	blog.Type = cnst.BlogTypeConstant.Me

	if err := biz.service.CheckUserExists(ctx, dto.UserId); err != nil {
		return err
	}

	meInfo, err := biz.service.FindDetailsBlog(ctx, map[string]interface{}{"user_id": dto.UserId, "type": cnst.BlogTypeConstant.Me})
	if err != nil {
		err = biz.service.Create(ctx, &blog)
		if err != nil {
			return common.ErrInternal(err)
		}
	} else {
		blog.Id = meInfo.Id
		updateData := bson.M{"$set": bson.M{
			"title":      blog.Title,
			"content":    blog.Content,
			"thumbnail":  blog.Thumbnail,
			"summary":    blog.Summary,
			"status":     blog.Status,
			"tag_ids":    blog.TagIds,
			"updated_at": &now,
		}}
		err = biz.service.Update(ctx, map[string]interface{}{"_id": meInfo.Id}, updateData)
		if err != nil {
			return common.ErrInternal(err)
		}
	}

	return nil
}
