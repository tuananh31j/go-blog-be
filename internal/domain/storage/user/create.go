package userStorage

import (
	"context"

	userModel "nta-blog/internal/domain/model/user"
)

func (s *store) CreateUser(ctx context.Context, dto *userModel.User) error {
	col := s.db.Collection(userModel.UserCollectionName)

	_, err := col.InsertOne(ctx, &dto)

	return err
}
