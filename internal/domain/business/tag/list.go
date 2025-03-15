package tagBusiness

import (
	"context"

	tagModel "nta-blog/internal/domain/model/tag"
)

type ListTagService interface {
	GetAllTag(ctx context.Context) ([]tagModel.Tag, error)
}

type listTagBiz struct {
	service ListTagService
}

func NewListTagBiz(sv ListTagService) *listTagBiz {
	return &listTagBiz{service: sv}
}

func (biz *listTagBiz) GetAllTag(ctx context.Context) ([]tagModel.Tag, error) {
	return biz.service.GetAllTag(ctx)
}
