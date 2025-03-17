package authService

import "context"

type RefreshTokenStore interface {
	SaveRefreshToken(ctx context.Context, token, userId string) error
	RemoveToken(ctx context.Context, userId string) error
	FindToken(ctx context.Context, userId string) (string, error)
}

type refreshTokenService struct {
	store RefreshTokenStore
}

func NewRefreshTokenService(store RefreshTokenStore) *refreshTokenService {
	return &refreshTokenService{store: store}
}

func (s *refreshTokenService) SaveRefreshToken(ctx context.Context, token, userId string) error {
	return s.store.SaveRefreshToken(ctx, token, userId)
}

func (s *refreshTokenService) RemoveRefreshToken(ctx context.Context, userId string) error {
	return s.store.RemoveToken(ctx, userId)
}

func (s *refreshTokenService) CheckRefreshTokenExists(ctx context.Context, userId string) (string, error) {
	return s.store.FindToken(ctx, userId)
}
