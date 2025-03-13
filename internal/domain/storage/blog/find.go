package blogStorage

import (
	"context"

	blogModel "nta-blog/internal/domain/model/blog"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) Find(ctx context.Context, conditions map[string]interface{}) (*blogModel.Blog, error) {
	col := s.db.Collection(blogModel.BlogCollection)
	var result blogModel.Blog
	filter := bson.M(conditions)
	err := col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
