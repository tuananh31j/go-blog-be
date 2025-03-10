package authtore

import (
	"context"
	"encoding/json"
	"fmt"

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

	// Chuyển struct thành JSON string
	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// In ra string
	fmt.Println(string(userJSON))
	return &user, nil
}
