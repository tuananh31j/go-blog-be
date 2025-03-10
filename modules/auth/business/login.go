package authBusiness

import (
	"context"
	"fmt"

	"nta-blog/common"
	authmdl "nta-blog/modules/auth/model"

	"github.com/rs/zerolog"
)

type LoginRepo interface {
	GetUserByEmail(ctx context.Context, email string) (*authmdl.User, error)
	SaveRefreshToken(ctx context.Context, token, userId string) error
}

type Hasher interface {
	Hash() string
	SetSalt(s string)
}

type loginBiz struct {
	repo   LoginRepo
	hasher Hasher
	logger *zerolog.Logger
}

func NewLoginBiz(repo LoginRepo, hashser Hasher) *loginBiz {
	return &loginBiz{repo: repo, hasher: hashser}
}

func (biz *loginBiz) Login(ctx context.Context, data authmdl.LoginDTO) (accessToken string, refreshToken string, err error) {
	user, err := biz.repo.GetUserByEmail(ctx, data.Email)
	if err != nil {
		biz.logger.Debug().Msg(fmt.Sprintf("Recover>>>>>>> %v", err))
		return "", "", common.NewErrorResponse(err, "Not valid!", err.Error())
	}
	biz.hasher.SetSalt(user.Salt)
	hash := biz.hasher.Hash()
	if hash != user.Password {
		return "", "", common.NewCustomError(err, "Your data is not valid!", "Password is wrong!")
	}
	accessToken = user.CreateAccessToken()
	refreshToken = user.CreateRefreshToken()

	return accessToken, refreshToken, nil
}
