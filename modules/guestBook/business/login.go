package authBusiness

import (
	"context"
	"fmt"

	"nta-blog/common"
	"nta-blog/libs/hashser"
	authmdl "nta-blog/modules/auth/model"

	"github.com/rs/zerolog"
)

type LoginRepo interface {
	GetUserByEmail(ctx context.Context, email string) (*authmdl.User, error)
	SaveRefreshToken(ctx context.Context, token, userId string) error
}

type loginBiz struct {
	repo   LoginRepo
	logger *zerolog.Logger
}

func NewLoginBiz(repo LoginRepo, log *zerolog.Logger) *loginBiz {
	return &loginBiz{repo: repo, logger: log}
}

func (biz *loginBiz) Login(ctx context.Context, data authmdl.LoginDTO) (accessToken string, refreshToken string, err error) {
	user, err := biz.repo.GetUserByEmail(ctx, data.Email)
	if err != nil {
		biz.logger.Debug().Msg(fmt.Sprintf("Recover>>>>>>> %v", err))
		return "", "", common.NewErrorResponse(err, "Not valid!", err.Error())
	}
	hash := hashser.Hash(data.Password, user.Salt)
	if hash != user.Password {
		return "", "", common.NewCustomError(err, "Your data is not valid!", "Password is wrong!")
	}
	accessToken = user.CreateAccessToken()
	refreshToken = user.CreateRefreshToken()

	go func() {
		defer common.AppRecover()
		if err := biz.repo.SaveRefreshToken(ctx, refreshToken, user.Id.Hex()); err != nil {
			panic(common.ErrSideEffectSaveRefreshToken(err))
		}
	}()

	return accessToken, refreshToken, nil
}
