package guestBookStorage

import (
	"context"

	guestbookModel "nta-blog/internal/domain/model/guestBook"
	"nta-blog/internal/lib/logger"
)

func (s *store) Create(ctx context.Context, dto *guestbookModel.GuestBook) error {
	col := s.db.Collection(guestbookModel.GuestBookCollection)
	_, err := col.InsertOne(ctx, &dto)
	if err != nil {
		logger.ZeroLog.Debug().Interface("dto", dto).Msg("Guest book data before insert")

		return err
	}
	return nil
}
