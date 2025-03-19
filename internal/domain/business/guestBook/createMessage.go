package guestbookBusiness

import (
	"context"
	"time"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	guestbookModel "nta-blog/internal/domain/model/guestBook"
	userModel "nta-blog/internal/domain/model/user"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateGuestBookService interface {
	FindOneUser(ctx context.Context, userId primitive.ObjectID) (*userModel.User, error)
	CreateMessage(ctx context.Context, dto *guestbookModel.GuestBook) error
}

type createGuestBookBiz struct {
	service CreateGuestBookService
	logger  *zerolog.Logger
}

func NewCreateGuestBookBiz(sv CreateGuestBookService, log *zerolog.Logger) *createGuestBookBiz {
	return &createGuestBookBiz{service: sv, logger: log}
}

func (biz *createGuestBookBiz) CreateMessage(ctx context.Context, msg string, userId primitive.ObjectID) error {
	var message guestbookModel.GuestBook
	message.Message = msg
	message.UserId = userId
	message.Status = cnst.StatusMessage.Actived
	now := time.Now()
	message.CreatedAt = &now
	message.UpdatedAt = &now
	message.Id = primitive.NewObjectID()
	user, err := biz.service.FindOneUser(ctx, userId)
	if err != nil {
		biz.logger.Debug().Interface("userId", userId).Msg("Failed to find user")
		return common.ErrBadRequest(err)
	}

	if user.Status == cnst.StatusAccount.Banned {
		biz.logger.Debug().Interface("userId", userId).Msg("User is banned")
		return common.NewErrorResponse(err, "Bạn đã bị cấm dùng chức năng này!", "User is banned!")
	}

	if err := biz.service.CreateMessage(ctx, &message); err != nil {
		biz.logger.Debug().Err(err).Msg("Failed to create message")
		return common.ErrInternal(err)
	}
	return nil
}
