package tagStorage

import (
	"context"

	tagModel "nta-blog/internal/domain/model/tag"
)

func (s *store) Create(ctx context.Context, dto *tagModel.Tag) error {
	col := s.db.Collection(tagModel.TagCollection)
	_, err := col.InsertOne(ctx, &dto)
	return err
}
