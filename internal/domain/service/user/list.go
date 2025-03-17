package userService

import (
	"context"

	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson"
)

type ListUserStore interface {
	List(ctx context.Context, pipeline bson.A) ([]userModel.User, error)
}

type listUserService struct {
	store ListUserStore
}

func NewListUserService(store ListUserStore) *listUserService {
	return &listUserService{store: store}
}

func (s *listUserService) ListUser(ctx context.Context, pipeline bson.A) ([]userModel.User, error) {
	return s.store.List(ctx, pipeline)
}
