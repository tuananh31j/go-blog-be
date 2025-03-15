package userModel

import (
	"fmt"
	"time"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	"nta-blog/internal/infrastructure/config"
)

const UserCollectionName = "users"

type User struct {
	common.CommonModal `bson:",inline"`
	Name               string              `bson:"name" json:"name"`
	Password           string              `bson:"password" json:"omitempty"`
	Salt               string              `bson:"salt" json:"salt"`
	Email              string              `bson:"email" json:"email"`
	Role               cnst.TRoleAccount   `bson:"role" json:"role"`
	Status             cnst.TStatusAccount `bson:"status" json:"status"`
}

func (u *User) CreateAccessToken() string {
	exp := time.Now().Add(time.Minute * 10000000).Unix()
	token := common.GenerateJWT(config.Env.SecretAccessKey, map[string]string{"id": u.Id.Hex(), "role": fmt.Sprintf("%v", u.Role)}, exp)

	return token
}

func (u *User) CreateRefreshToken() string {
	exp := time.Now().Add(time.Hour * 24 * 7).Unix()
	token := common.GenerateJWT(config.Env.SecretRefreshKey, map[string]string{"id": u.Id.Hex(), "role": string(u.Role)}, exp)

	return token
}

func (u *User) VerifyToken(tokenString string, secret string) (*common.Payload, error) {
	return common.VerifyJWT(tokenString, secret)
}
