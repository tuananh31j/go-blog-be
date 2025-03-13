package guestBookStorage

import (
	"context"

	guestbookModel "nta-blog/internal/domain/model/guestBook"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) List(ctx context.Context, pipeline bson.A) ([]guestbookModel.GuestBook, error) {
	var result []guestbookModel.GuestBook
	col := s.db.Collection(guestbookModel.GuestBookCollection)
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
