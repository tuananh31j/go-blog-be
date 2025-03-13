package guestbookService

import (
	"context"

	guestbookModel "nta-blog/internal/domain/model/guestBook"
	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateGuestBookStore interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
	Create(ctx context.Context, dto *guestbookModel.GuestBook) error
}

type createGuestBookService struct {
	store CreateGuestBookStore
}

func NewCreateGuestBookService(s CreateGuestBookStore) *createGuestBookService {
	return &createGuestBookService{store: s}
}

func (sv *createGuestBookService) FindOneUser(ctx context.Context, userId primitive.ObjectID) (*userModel.User, error) {
	return sv.store.FindOneUser(ctx, map[string]interface{}{"_id": userId})
}

func (sv *createGuestBookService) CreateMessage(ctx context.Context, dto *guestbookModel.GuestBook) error {
	return sv.store.Create(ctx, dto)
}
