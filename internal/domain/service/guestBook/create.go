package guestbookService

import (
	"context"

	guestbookModel "nta-blog/internal/domain/model/guestBook"
	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateGuestBookStore interface {
	Create(ctx context.Context, dto *guestbookModel.GuestBook) error
}

type CheckUserStore interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
}

type createGuestBookService struct {
	guestBookStore CreateGuestBookStore
	userStore      CheckUserStore
}

func NewCreateGuestBookService(gbs CreateGuestBookStore, us CheckUserStore) *createGuestBookService {
	return &createGuestBookService{guestBookStore: gbs, userStore: us}
}

func (sv *createGuestBookService) FindOneUser(ctx context.Context, userId primitive.ObjectID) (*userModel.User, error) {
	return sv.userStore.FindOneUser(ctx, map[string]interface{}{"_id": userId})
}

func (sv *createGuestBookService) CreateMessage(ctx context.Context, dto *guestbookModel.GuestBook) error {
	return sv.guestBookStore.Create(ctx, dto)
}
