package tagBusiness

import (
	"context"
	"fmt"

	"nta-blog/internal/common"
	tagModel "nta-blog/internal/domain/model/tag"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateTagService interface {
	CheckTagNameExists(ctx context.Context, tagName string) error
	CheckTagIdExists(ctx context.Context, tagId primitive.ObjectID) (tagModel.Tag, error)
	UpdateTag(ctx context.Context, tagId primitive.ObjectID, tagName string) error
}

type updateTagBiz struct {
	service UpdateTagService
}

func NewUpdateTagBiz(sv UpdateTagService) *updateTagBiz {
	return &updateTagBiz{service: sv}
}

func (biz *updateTagBiz) UpdateTagBiz(ctx context.Context, tagId primitive.ObjectID, tagName string) error {
	currentTag, err := biz.service.CheckTagIdExists(ctx, tagId)
	if err != nil {
		return common.NewErrorResponse(err, "Tag is not exists!", fmt.Sprintf("This id %v is not exists!", tagId.Hex()))
	}

	if err := biz.service.CheckTagNameExists(ctx, tagName); err == nil && currentTag.Id != tagId {
		return common.NewErrorResponse(err, "This name already exists!", "This name already exists!")
	}
	if err := biz.service.UpdateTag(ctx, tagId, tagName); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
