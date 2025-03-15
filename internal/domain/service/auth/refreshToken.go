package authService

import "context"

type RefreshTokenStore interface {
	SaveRefreshToken(ctx context.Context, token string) error
	RemoveToken(ctx context.Context, token string) error
	FindToken(ctx context.Context, token string) (string, error)
}

type refreshTokenService struct {
	store RefreshTokenStore
}

func NewRefreshTokenService(store RefreshTokenStore) *refreshTokenService {
	return &refreshTokenService{store: store}
}

func (s *refreshTokenService) SaveRefreshToken(ctx context.Context, token string) error {
	return s.store.SaveRefreshToken(ctx, token)
}

func (s *refreshTokenService) RemoveRefreshToken(ctx context.Context, token string) error {
	return s.store.RemoveToken(ctx, token)
}

func (s *refreshTokenService) CheckRefeshTokenExists(ctx context.Context, token string) (string, error) {
	return s.store.FindToken(ctx, token)
}
