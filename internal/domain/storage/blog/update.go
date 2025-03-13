package blogStorage

import (
	"context"

	blogModel "nta-blog/internal/domain/model/blog"
)

func (s *store) Update(ctx context.Context, dto *blogModel.Blog) error {
	col := s.db.Collection(blogModel.BlogCollection)
	_, err := col.InsertOne(ctx, &dto)
	if err != nil {
		return err
	}
	return nil
}
