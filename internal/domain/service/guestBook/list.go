package guestbookService

import (
	"context"

	guestbookModel "nta-blog/internal/domain/model/guestBook"

	"go.mongodb.org/mongo-driver/bson"
)

type ListGuestBookStore interface {
	ListMessage(ctx context.Context, pipeline bson.A) ([]guestbookModel.GuestBook, error)
}

type listGuestBookService struct {
	store ListGuestBookStore
}

func NewListGuestBookService(s ListGuestBookStore) *listGuestBookService {
	return &listGuestBookService{store: s}
}

func (sv *listGuestBookService) GetListMessage(ctx context.Context, pipeline bson.A) ([]guestbookModel.GuestBook, error) {
	return sv.store.ListMessage(ctx, pipeline)
}
