package userStorage

import (
	"context"

	cnst "nta-blog/internal/constant"
	userModel "nta-blog/internal/domain/model/user"
)

func (s *store) CreateUser(ctx context.Context, dto *userModel.User) error {
	col := s.db.Collection(cnst.UserCollection)

	_, err := col.InsertOne(ctx, &dto)

	return err
}
