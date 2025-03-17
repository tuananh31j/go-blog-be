package userService

import (
	"context"

	cnst "nta-blog/internal/constant"
	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type banUserStore interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) error
}

type banUserService struct {
	store banUserStore
}

func NewBanUserService(store banUserStore) *banUserService {
	return &banUserService{store: store}
}

func (s *banUserService) BanUser(ctx context.Context, userId primitive.ObjectID) error {
	conditions := map[string]interface{}{"_id": userId}
	dataFiled := map[string]interface{}{"status": cnst.StatusAccount.Banned}
	err := s.store.UpdateUser(ctx, conditions, dataFiled)
	if err != nil {
		return err
	}

	return nil
}

func (s *banUserService) CheckUserExists(ctx context.Context, userId primitive.ObjectID) (*userModel.User, error) {
	conditions := map[string]interface{}{"_id": userId}
	user, err := s.store.FindOneUser(ctx, conditions)
	if err != nil {
		return nil, err
	}

	return user, nil
}
