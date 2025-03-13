package tagBusiness

import (
	"context"

	"nta-blog/internal/common"
	tagModel "nta-blog/internal/domain/model/tag"
)

type CreateTagService interface {
	CheckTagExists(ctx context.Context, tagName string) error
	CreateNewTag(ctx context.Context, tag tagModel.TagDTO) error
}

type createTagBiz struct {
	service CreateTagService
}

func NewCreateTagBiz(sv CreateTagService) *createTagBiz {
	return &createTagBiz{service: sv}
}

func (biz *createTagBiz) CreateNewTag(ctx context.Context, tag tagModel.TagDTO) error {
	if err := biz.service.CheckTagExists(ctx, tag.Name); err == nil {
		return common.NewErrorResponse(err, "Tên thẻ đã tồn tại", "Tag name is exists!")
	}

	if err := biz.service.CreateNewTag(ctx, tag); err != nil {
		return common.ErrInternal(err)
	}

	return nil
}
