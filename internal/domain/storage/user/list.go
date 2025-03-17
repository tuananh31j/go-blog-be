package userStorage

import (
	"context"

	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) List(ctx context.Context, pipeline bson.A) ([]userModel.User, error) {
	var users []userModel.User

	userCol := s.db.Collection(userModel.UserCollectionName)
	cursor, err := userCol.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
