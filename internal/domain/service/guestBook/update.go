package guestbookService

import (
	"context"

	guestbookModel "nta-blog/internal/domain/model/guestBook"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateGuestBookStore interface {
	Find(ctx context.Context, condition map[string]interface{}) ([]*guestbookModel.GuestBook, error)
	Update(ctx context.Context, guestBookId primitive.ObjectID, updateField map[string]interface{}) error
}

type updateGuestBookService struct {
	guestBookStore UpdateGuestBookStore
}

func NewUpdateGuestBookStore(guestBookStore UpdateGuestBookStore) *updateGuestBookService {
	return &updateGuestBookService{guestBookStore: guestBookStore}
}

func (sv *updateGuestBookService) UpdateMessage(ctx context.Context, guestBookId primitive.ObjectID, updateField map[string]interface{}) error {
	return sv.guestBookStore.Update(ctx, guestBookId, updateField)
}

func (sv *updateGuestBookService) FindMessage(ctx context.Context, condition map[string]interface{}) ([]*guestbookModel.GuestBook, error) {
	return sv.guestBookStore.Find(ctx, condition)
}
