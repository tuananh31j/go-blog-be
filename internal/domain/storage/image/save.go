package imageStorage

import (
	"context"

	imageModel "nta-blog/internal/domain/model/image"
)

func (s *store) SaveImage(ctx context.Context, image *imageModel.Image) error {
	col := s.db.Collection(imageModel.ImageCollection)
	_, err := col.InsertOne(ctx, image)
	if err != nil {
		return err
	}
	return nil
}
