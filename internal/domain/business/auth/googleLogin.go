package authBusiness

import (
	"context"
	"time"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	userModel "nta-blog/internal/domain/model/user"
	"nta-blog/internal/lib/hashser"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GoogleSevice interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
	SaveRefreshToken(ctx context.Context, token, userId string) error
	CreateUser(ctx context.Context, dto *userModel.User) error
}

type googleLoginBiz struct {
	sevice GoogleSevice
	logger *zerolog.Logger
}

func NewGoogleLoginBiz(sevice GoogleSevice, log *zerolog.Logger) *googleLoginBiz {
	return &googleLoginBiz{sevice: sevice, logger: log}
}

func (biz *googleLoginBiz) GoogleLogin(ctx context.Context, data userModel.GoogleLoginDTO) (accessToken string, refreshToken string, err error) {
	user, err := biz.sevice.FindOneUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {

		passRand := common.GenSalt()
		salt := common.GenSalt()

		hash := hashser.Hash(passRand, salt)
		now := time.Now()
		newUser := userModel.User{
			CommonModal: common.CommonModal{
				Id:        primitive.NewObjectID(),
				CreatedAt: &now,
				UpdatedAt: &now,
			},
			Role:     cnst.UserRole,
			Name:     data.Name,
			Email:    data.Email,
			Password: hash,
			Salt:     salt,
		}
		if err := biz.sevice.CreateUser(ctx, &newUser); err != nil {
			return "", "", common.ErrInternal(err)
		}
		accessToken = newUser.CreateAccessToken()
		refreshToken = newUser.CreateRefreshToken()

		go func() {
			defer common.AppRecover()
			if err := biz.sevice.SaveRefreshToken(ctx, refreshToken, user.Id.Hex()); err != nil {
				panic(common.ErrSideEffectSaveRefreshToken(err))
			}
		}()

		return accessToken, refreshToken, nil
	}
	accessToken = user.CreateAccessToken()
	refreshToken = user.CreateRefreshToken()

	go func() {
		defer common.AppRecover()
		if err := biz.sevice.SaveRefreshToken(ctx, refreshToken, user.Id.Hex()); err != nil {
			panic(common.ErrSideEffectSaveRefreshToken(err))
		}
	}()

	return accessToken, refreshToken, nil
}
