package userBusiness

import (
	"context"

	"nta-blog/internal/common"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BanUserService interface {
	CheckUserExists(ctx context.Context, userId primitive.ObjectID) error
	BanUser(ctx context.Context, userId primitive.ObjectID) error
}

type banUserBiz struct {
	service BanUserService
}

func NewBanUserBiz(sv BanUserService) *banUserBiz {
	return &banUserBiz{service: sv}
}

func (biz *banUserBiz) BanUser(ctx context.Context, userId primitive.ObjectID) error {
	if err := biz.service.CheckUserExists(ctx, userId); err != nil {
		return common.ErrBadRequest(err)
	}

	if err := biz.service.BanUser(ctx, userId); err != nil {
		return common.ErrBadRequest(err)
	}
	return nil
}
