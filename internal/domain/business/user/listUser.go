package userBusiness

import (
	"context"

	"nta-blog/internal/common"
	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson"
)

type ListUserService interface {
	ListUser(ctx context.Context, pipeline bson.A) ([]userModel.User, error)
}

type listUserBiz struct {
	service ListUserService
}

func NewListUserBiz(sv ListUserService) *listUserBiz {
	return &listUserBiz{service: sv}
}

func (biz *listUserBiz) ListUserBiz(ctx context.Context, pipeLine bson.A) ([]userModel.User, error) {
	results, err := biz.service.ListUser(ctx, pipeLine)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return results, nil
}
