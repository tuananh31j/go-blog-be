package userModel

import (
	"time"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	"nta-blog/internal/infrastructure/config"
)

const UserCollectionName = "users"

type User struct {
	common.CommonModal
	Name     string              `bson:"name" json:"name"`
	Password string              `bson:"password" json:"omitempty"`
	Salt     string              `bson:"salt" json:"salt"`
	Email    string              `bson:"email" json:"email"`
	Role     cnst.TRoleAccount   `bson:"role" json:"role"`
	Status   cnst.TStatusAccount `bson:"status" json:"status"`
}

func (u *User) CreateAccessToken() string {
	exp := time.Now().Add(time.Minute * 15).Unix()
	token := common.GenerateJWT(config.Env.SecretAccessKey, map[string]string{"name": u.Name, "id": u.Id.Hex()}, exp)

	return token
}

func (u *User) CreateRefreshToken() string {
	exp := time.Now().Add(time.Hour * 24 * 7).Unix()
	token := common.GenerateJWT(config.Env.SecretRefreshKey, map[string]string{"name": u.Name, "id": u.Id.Hex()}, exp)

	return token
}
