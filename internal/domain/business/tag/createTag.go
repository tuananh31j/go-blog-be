package tagBusiness

import (
	"context"
	"time"

	"nta-blog/internal/common"
	tagModel "nta-blog/internal/domain/model/tag"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTagService interface {
	CheckTagExists(ctx context.Context, tagName string) error
	CreateNewTag(ctx context.Context, tag *tagModel.Tag) error
}

type createTagBiz struct {
	service CreateTagService
}

func NewCreateTagBiz(sv CreateTagService) *createTagBiz {
	return &createTagBiz{service: sv}
}

func (biz *createTagBiz) CreateNewTag(ctx context.Context, dto *tagModel.TagDTO) error {
	var tag tagModel.Tag
	now := time.Now()
	tag.CreatedAt = &now
	tag.UpdatedAt = &now
	tag.Name = dto.Name
	tag.Id = primitive.NewObjectID()

	if err := biz.service.CheckTagExists(ctx, tag.Name); err == nil {
		return common.NewErrorResponse(err, "Tên thẻ đã tồn tại", "Tag name is exists!")
	}

	if err := biz.service.CreateNewTag(ctx, &tag); err != nil {
		return common.ErrInternal(err)
	}

	return nil
}
