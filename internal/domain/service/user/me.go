package userService

import (
	"context"

	userModel "nta-blog/internal/domain/model/user"
)

type MeStore interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
}

type meService struct {
	store MeStore
}

func NewMeService(store MeStore) *meService {
	return &meService{store: store}
}

func (s *meService) GetMe(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error) {
	return s.store.FindOneUser(ctx, conditions)
}
