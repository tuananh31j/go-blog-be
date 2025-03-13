package guestBookStorage

import (
	"context"

	guestbookModel "nta-blog/internal/domain/model/guestBook"
)

func (s *store) Create(ctx context.Context, dto *guestbookModel.GuestBook) error {
	col := s.db.Collection(guestbookModel.GuestBookCollection)
	_, err := col.InsertOne(ctx, &dto)
	if err != nil {
		return err
	}
	return nil
}
