package imageBusiness

import (
	"context"

	imageModel "nta-blog/internal/domain/model/image"
)

type ListImagesService interface {
	GetListImage(ctx context.Context) ([]imageModel.Image, error)
}

type listImagesBiz struct {
	service ListImagesService
}

func NewListImagesBiz(sv ListImagesService) *listImagesBiz {
	return &listImagesBiz{service: sv}
}

func (biz *listImagesBiz) GetListImage(ctx context.Context) ([]imageModel.Image, error) {
	return biz.service.GetListImage(ctx)
}
