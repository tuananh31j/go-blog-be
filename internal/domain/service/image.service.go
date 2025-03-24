package service

import (
	"context"
	"mime/multipart"

	repository "nta-blog/internal/domain/interface"
	imageModel "nta-blog/internal/domain/model/image"
)

type imageService struct {
	store repository.ImageRepository
}

func NewImageService(store repository.ImageRepository) *imageService {
	return &imageService{store: store}
}

func (sv *imageService) GetListImage(ctx context.Context) ([]imageModel.Image, error) {
	return sv.store.ListImages(ctx)
}

func (s *imageService) UploadFile(ctx context.Context, file multipart.File) (*imageModel.UploadResFormCld, error) {
	return s.store.UploadFile(ctx, file)
}

func (s *imageService) SaveImage(ctx context.Context, image *imageModel.Image) error {
	return s.store.SaveImage(ctx, image)
}
