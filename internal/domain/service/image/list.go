package imageService

import (
	"context"

	imageModel "nta-blog/internal/domain/model/image"
)

type ListImageStore interface {
	ListImages(ctx context.Context) ([]imageModel.Image, error)
}

type listImageService struct {
	store ListImageStore
}

func NewListImageService(s ListImageStore) *listImageService {
	return &listImageService{store: s}
}

func (sv *listImageService) GetListImage(ctx context.Context) ([]imageModel.Image, error) {
	return sv.store.ListImages(ctx)
}
