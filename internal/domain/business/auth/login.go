package authBusiness

import (
	"context"
	"encoding/json"
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

func (biz *loginBiz) Login(ctx context.Context, data userModel.LoginDTO) (accessToken string, refreshToken string, userTiny *userModel.PrivateUser, err error) {
	user, err := biz.service.FindOneUser(ctx, map[string]interface{}{"email": data.Email})
	jsonData, _ := json.Marshal(user)
	biz.logger.Info().Interface("user", string(jsonData)).Msg("User data")
	if err != nil {
		biz.logger.Debug().Msg(fmt.Sprintf("Recover>>>>>>> %v", err))
		return "", "", nil, common.NewErrorResponse(err, "Not valid!", err.Error())
	}
	hash := hashser.Hash(data.Password, user.Salt)
	if hash != user.Password {
		return "", "", nil, common.NewCustomError(err, "Your data is not valid!", "Password is wrong!")
	}
	accessToken = user.CreateAccessToken()
	refreshToken = user.CreateRefreshToken()
	userTiny = &userModel.PrivateUser{
		Id:       user.Id.Hex(),
		NameFake: user.NameFake,
		Email:    user.Email,
		Role:     user.Role,
		Avt:      user.Avt,
	}

	go func() {
		defer common.AppRecover()
		if err := biz.service.SaveRefreshToken(ctx, refreshToken, user.Id.Hex()); err != nil {
			panic(common.ErrSideEffectSaveRefreshToken(err))
		}
	}()

	return accessToken, refreshToken, userTiny, nil
}
