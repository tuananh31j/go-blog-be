package service

import (
	"context"

	cnst "nta-blog/internal/constant"
	repository "nta-blog/internal/domain/interface"
	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct {
	userStore repository.UserRepository
}

func NewUserService(us repository.UserRepository) *userService {
	return &userService{userStore: us}
}

func (s *userService) BanUser(ctx context.Context, userId primitive.ObjectID) error {
	conditions := map[string]interface{}{"_id": userId}
	dataFiled := map[string]interface{}{"status": cnst.StatusAccount.Banned}
	err := s.userStore.UpdateUser(ctx, conditions, dataFiled)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) CheckUserExists(ctx context.Context, userId primitive.ObjectID) (*userModel.User, error) {
	conditions := map[string]interface{}{"_id": userId}
	user, err := s.userStore.FindOneUser(ctx, conditions)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) ListUser(ctx context.Context, pipeline bson.A) ([]userModel.User, error) {
	return s.userStore.List(ctx, pipeline)
}

func (s *userService) GetMe(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error) {
	return s.userStore.FindOneUser(ctx, conditions)
}
