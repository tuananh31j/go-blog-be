package imageStorage

import (
	"context"

	imageModel "nta-blog/internal/domain/model/image"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) ListImages(ctx context.Context) ([]imageModel.Image, error) {
	col := s.db.Collection(imageModel.ImageCollection)
	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var images []imageModel.Image
	if err = cursor.All(ctx, &images); err != nil {
		return nil, err
	}
	return images, nil
}
