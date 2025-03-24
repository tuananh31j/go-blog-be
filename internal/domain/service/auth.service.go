package service

import (
	"context"

	repository "nta-blog/internal/domain/interface"
	userModel "nta-blog/internal/domain/model/user"
)

type authService struct {
	userstore  repository.UserRepository
	tokenStore repository.TokenRepository
}

func NewAuthService(us repository.UserRepository, ts repository.TokenRepository) *authService {
	return &authService{userstore: us, tokenStore: ts}
}

func (s *authService) CheckRefreshTokenExists(ctx context.Context, userId string) (string, error) {
	return s.tokenStore.FindToken(ctx, userId)
}

func (s *authService) RemoveRefreshToken(ctx context.Context, userId string) error {
	return s.tokenStore.RemoveToken(ctx, userId)
}

func (sv *authService) CreateUser(ctx context.Context, dto *userModel.User) error {
	return sv.userstore.CreateUser(ctx, dto)
}

func (sv *authService) SaveRefreshToken(ctx context.Context, token, userId string) error {
	return sv.tokenStore.SaveRefreshToken(ctx, token, userId)
}

func (sv *authService) FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error) {
	return sv.userstore.FindOneUser(
		ctx, conditions,
	)
}

func (sv *authService) UpdateUser(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) error {
	return sv.userstore.UpdateUser(ctx, filter, data)
}

func (sv *authService) GetToken(ctx context.Context, userId string) (string, error) {
	return sv.tokenStore.FindToken(ctx, userId)
}
