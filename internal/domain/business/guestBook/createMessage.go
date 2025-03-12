package guestbookBusiness

import (
	"context"

	"nta-blog/internal/common"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateGuestBookService interface {
	FindOneUser(ctx context.Context, userId primitive.ObjectID) error
	CreateMessage(ctx context.Context, msg string, userId primitive.ObjectID) error
}

type createGuestBookBiz struct {
	service CreateGuestBookService
}

func NewGuestBookBiz(sv CreateGuestBookService) *createGuestBookBiz {
	return &createGuestBookBiz{service: sv}
}

func (biz *createGuestBookBiz) CreateMessage(ctx context.Context, msg string, userId primitive.ObjectID) error {
	if err := biz.service.FindOneUser(ctx, userId); err != nil {
		return common.ErrBadRequest(err)
	}

	if err := biz.service.CreateMessage(ctx, msg, userId); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
