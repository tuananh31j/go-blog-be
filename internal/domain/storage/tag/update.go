package tagStorage

import (
	"context"

	tagModel "nta-blog/internal/domain/model/tag"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) Update(ctx context.Context, condition map[string]interface{}, dto map[string]interface{}) error {
	col := s.db.Collection(tagModel.TagCollection)
	filter := bson.M(condition)
	updateField := bson.M{"$set": bson.M(dto)}
	_, err := col.UpdateOne(ctx, filter, updateField)
	return err
}
