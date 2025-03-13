package guestBookStorage

import (
	"context"

	guestbookModel "nta-blog/internal/domain/model/guestBook"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) Find(ctx context.Context, condition map[string]interface{}) ([]*guestbookModel.GuestBook, error) {
	col := s.db.Collection(guestbookModel.GuestBookCollection)
	filter := bson.M(condition)
	cur, err := col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var res []*guestbookModel.GuestBook
	for cur.Next(ctx) {
		var item guestbookModel.GuestBook
		err := cur.Decode(&item)
		if err != nil {
			return nil, err
		}
		res = append(res, &item)
	}
	return res, nil
}
