package imageService

import (
	"context"
	"mime/multipart"

	imageModel "nta-blog/internal/domain/model/image"
)

type UploadStore interface {
	UploadFile(ctx context.Context, file multipart.File) (*imageModel.UploadResFormCld, error)
	SaveImage(ctx context.Context, image *imageModel.Image) error
}

type imageService struct {
	store UploadStore
}

func NewUploadService(store UploadStore) *imageService {
	return &imageService{store: store}
}

func (s *imageService) UploadFile(ctx context.Context, file multipart.File) (*imageModel.UploadResFormCld, error) {
	return s.store.UploadFile(ctx, file)
}

func (s *imageService) SaveImage(ctx context.Context, image *imageModel.Image) error {
	return s.store.SaveImage(ctx, image)
}
