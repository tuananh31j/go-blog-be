package guestbookBusiness

import (
	"context"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	guestbookModel "nta-blog/internal/domain/model/guestBook"
	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateGuestBookService interface {
	FindOneUser(ctx context.Context, userId primitive.ObjectID) (*userModel.User, error)
	CreateMessage(ctx context.Context, dto *guestbookModel.GuestBook) error
}

type createGuestBookBiz struct {
	service CreateGuestBookService
}

func NewGuestBookBiz(sv CreateGuestBookService) *createGuestBookBiz {
	return &createGuestBookBiz{service: sv}
}

func (biz *createGuestBookBiz) CreateMessage(ctx context.Context, msg string, userId primitive.ObjectID) error {
	var message guestbookModel.GuestBook
	message.Message = msg
	message.UserId = userId
	message.Status = cnst.StatusMessage.Pending
	user, err := biz.service.FindOneUser(ctx, userId)
	if err != nil {
		return common.ErrBadRequest(err)
	}

	if user.Status == cnst.StatusAccount.Banned {
		return common.NewErrorResponse(err, "Bạn đã bị cấm dùng chức năng này!", "User is banned!")
	}

	if err := biz.service.CreateMessage(ctx, &message); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
