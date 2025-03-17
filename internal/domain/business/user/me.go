package userBusiness

import (
	"context"

	"nta-blog/internal/common"
	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MeService interface {
	GetMe(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
}

type meBiz struct {
	service MeService
}

func NewMeBiz(sv MeService) *meBiz {
	return &meBiz{service: sv}
}

func (biz *meBiz) GetMe(ctx context.Context, userId primitive.ObjectID) (*userModel.PrivateUser, error) {
	conditions := map[string]interface{}{"_id": userId}
	user, err := biz.service.GetMe(ctx, conditions)
	if err != nil {
		return nil, common.ErrBadRequest(err)
	}
	var userSimp userModel.PrivateUser = userModel.PrivateUser{
		Id:    user.Id.Hex(),
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Avt:   user.Avt,
	}

	return &userSimp, nil
}
