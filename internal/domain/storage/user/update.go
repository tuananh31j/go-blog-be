package userStorage

import (
	"context"

	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) UpdateUser(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) error {
	var user userModel.User
	condition := bson.M(filter)
	dataFiled := bson.M(data)
	userCol := s.db.Collection(userModel.UserCollectionName)
	err := userCol.FindOneAndUpdate(ctx, condition, bson.M{"$set": dataFiled}).Decode(&user)
	if err != nil {
		return err
	}

	return nil
}
