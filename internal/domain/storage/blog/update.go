package blogStorage

import (
	"context"

	blogModel "nta-blog/internal/domain/model/blog"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) Update(ctx context.Context, conditions map[string]interface{}, dto *blogModel.Blog) error {
	col := s.db.Collection(blogModel.BlogCollection)
	filter := bson.M(conditions)
	_, err := col.UpdateOne(ctx, filter, &dto)
	if err != nil {
		return err
	}
	return nil
}
