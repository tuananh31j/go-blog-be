package loginService

import (
	"context"

	userModel "nta-blog/internal/domain/model/user"
)

type UserStore interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
	CreateUser(ctx context.Context, dto *userModel.User) error
}
type TokenStore interface {
	SaveRefreshToken(ctx context.Context, token, userId string) error
}

type loginSevice struct {
	userstore  UserStore
	tokenStore TokenStore
}

func NewLoginService(us UserStore, ts TokenStore) *loginSevice {
	return &loginSevice{userstore: us, tokenStore: ts}
}

func (sv *loginSevice) FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error) {
	return sv.userstore.FindOneUser(
		ctx, conditions,
	)
}

func (sv *loginSevice) CreateUser(ctx context.Context, dto *userModel.User) error {
	return sv.userstore.CreateUser(ctx, dto)
}

func (sv *loginSevice) SaveRefreshToken(ctx context.Context, token, userId string) error {
	return sv.tokenStore.SaveRefreshToken(ctx, token, userId)
}
