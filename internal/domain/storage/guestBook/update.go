package guestBookStorage

import (
	"context"

	guestbookModel "nta-blog/internal/domain/model/guestBook"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *store) Update(ctx context.Context, guestBookId primitive.ObjectID, updateField map[string]interface{}) error {
	col := s.db.Collection(guestbookModel.GuestBookCollection)
	_, err := col.UpdateOne(ctx, bson.M{"_id": guestBookId}, bson.M{"$set": bson.M(updateField)})
	if err != nil {
		return err
	}
	return nil
}
