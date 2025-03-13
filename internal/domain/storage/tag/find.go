package tagStorage

import (
	"context"

	tagModel "nta-blog/internal/domain/model/tag"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) Find(ctx context.Context, conditions map[string]interface{}) (*tagModel.Tag, error) {
	var result tagModel.Tag
	col := s.db.Collection(tagModel.TagCollection)
	filter := bson.M(conditions)
	if err := col.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
