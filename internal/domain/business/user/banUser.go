package userBusiness

import (
	"context"

	"nta-blog/internal/common"
	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BanUserService interface {
	CheckUserExists(ctx context.Context, userId primitive.ObjectID) (*userModel.User, error)
	BanUser(ctx context.Context, userId primitive.ObjectID) error
}

type banUserBiz struct {
	service BanUserService
}

func NewBanUserBiz(sv BanUserService) *banUserBiz {
	return &banUserBiz{service: sv}
}

func (biz *banUserBiz) BanUser(ctx context.Context, userId primitive.ObjectID) error {
	_, err := biz.service.CheckUserExists(ctx, userId)
	if err != nil {
		return common.ErrBadRequest(err)
	}

	if err := biz.service.BanUser(ctx, userId); err != nil {
		return common.ErrBadRequest(err)
	}
	return nil
}
