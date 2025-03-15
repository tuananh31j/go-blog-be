package authService

import (
	"context"

	userModel "nta-blog/internal/domain/model/user"
)

type UserStore interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
	CreateUser(ctx context.Context, dto *userModel.User) error
}
type TokenStore interface {
	SaveRefreshToken(ctx context.Context, token string) error
}

type loginService struct {
	userstore  UserStore
	tokenStore TokenStore
}

func NewLoginService(us UserStore, ts TokenStore) *loginService {
	return &loginService{userstore: us, tokenStore: ts}
}

func (sv *loginService) FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error) {
	return sv.userstore.FindOneUser(
		ctx, conditions,
	)
}

func (sv *loginService) CreateUser(ctx context.Context, dto *userModel.User) error {
	return sv.userstore.CreateUser(ctx, dto)
}

func (sv *loginService) SaveRefreshToken(ctx context.Context, token string) error {
	return sv.tokenStore.SaveRefreshToken(ctx, token)
}
