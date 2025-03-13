package blogStorage

import (
	"context"

	blogModel "nta-blog/internal/domain/model/blog"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) List(ctx context.Context, pipeline bson.A) ([]blogModel.Blog, error) {
	col := s.db.Collection(blogModel.BlogCollection)
	var result []blogModel.Blog
	cursor, err := col.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
