package authService

import "context"

type LogoutStore interface {
	RemoveToken(ctx context.Context, userId string) error
}

type logoutService struct {
	store LogoutStore
}

func NewLogoutService(store LogoutStore) *logoutService {
	return &logoutService{store: store}
}

func (s *logoutService) RemoveRefreshToken(ctx context.Context, userId string) error {
	return s.store.RemoveToken(ctx, userId)
}
