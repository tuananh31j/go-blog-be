package authBusiness

import (
	"context"
	"fmt"

	"nta-blog/internal/common"
	userModel "nta-blog/internal/domain/model/user"
	"nta-blog/internal/lib/hashser"

	"github.com/rs/zerolog"
)

type LoginSevice interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
	SaveRefreshToken(ctx context.Context, token, userId string) error
}

type loginBiz struct {
	service LoginSevice
	logger  *zerolog.Logger
}

func NewLoginBiz(service LoginSevice, log *zerolog.Logger) *loginBiz {
	return &loginBiz{service: service, logger: log}
}

func (biz *loginBiz) Login(ctx context.Context, data userModel.LoginDTO) (accessToken string, refreshToken string, err error) {
	user, err := biz.service.FindOneUser(ctx, map[string]interface{}{"email": data.Email})
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
		if err := biz.service.SaveRefreshToken(ctx, refreshToken, user.Id.Hex()); err != nil {
			panic(common.ErrSideEffectSaveRefreshToken(err))
		}
	}()

	return accessToken, refreshToken, nil
}
