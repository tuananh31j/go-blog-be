package authBusiness

import (
	"context"

	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConfirmService interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) error
	GetToken(ctx context.Context, userId string) (string, error)
}

type confirmBiz struct {
	service ConfirmService
}

func NewConfirmBiz(sv ConfirmService) *confirmBiz {
	return &confirmBiz{service: sv}
}

func (biz *confirmBiz) Confirm(ctx context.Context, userId primitive.ObjectID, UserName string) (string, *userModel.PrivateUser, error) {
	conditions := map[string]interface{}{
		"_id": userId,
	}
	data := map[string]interface{}{
		"name_fake": UserName,
	}
	userstore, err := biz.service.FindOneUser(ctx, conditions)
	if err != nil {
		return "", nil, err
	}
	err = biz.service.UpdateUser(ctx, conditions, data)
	if err != nil {
		return "", nil, err
	}
	refreshToken, err := biz.service.GetToken(ctx, userId.Hex())
	if err != nil {
		return "", nil, err
	}
	userTiny := userModel.PrivateUser{
		Id:    userId.Hex(),
		Name:  UserName,
		Email: userstore.Email,
		Role:  userstore.Role,
		Avt:   userstore.Avt,
	}

	return refreshToken, &userTiny, nil
}
