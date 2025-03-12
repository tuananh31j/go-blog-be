package userStorage

import (
	"context"

	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error) {
	var user userModel.User
	filter := bson.M(conditions)
	userCol := s.db.Collection(userModel.UserCollectionName)
	err := userCol.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
