package tagStorage

import (
	"context"

	tagModel "nta-blog/internal/domain/model/tag"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) List(ctx context.Context, conditions map[string]interface{}) ([]tagModel.Tag, error) {
	var result []tagModel.Tag
	col := s.db.Collection(tagModel.TagCollection)
	filter := bson.M(conditions)
	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}
