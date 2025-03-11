package authtore

import (
	"context"

	cnst "nta-blog/constant"
	authmdl "nta-blog/modules/auth/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) FindOneUser(ctx context.Context, conditions map[string]interface{}) (*authmdl.User, error) {
	var user authmdl.User
	filter := bson.M(conditions)
	userCol := s.db.Collection(cnst.UserCollection)
	err := userCol.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
