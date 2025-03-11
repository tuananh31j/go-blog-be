package authBusiness

import (
	"context"
	"time"

	"nta-blog/common"
	cnst "nta-blog/constant"
	"nta-blog/libs/hashser"
	authmdl "nta-blog/modules/auth/model"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GoogleRepo interface {
	GetUserByEmail(ctx context.Context, email string) (*authmdl.User, error)
	SaveRefreshToken(ctx context.Context, token, userId string) error
	CreateUser(ctx context.Context, dto *authmdl.User) error
}

type googleLoginBiz struct {
	repo   GoogleRepo
	logger *zerolog.Logger
}

func NewGoogleLoginBiz(repo GoogleRepo, log *zerolog.Logger) *googleLoginBiz {
	return &googleLoginBiz{repo: repo, logger: log}
}

func (biz *googleLoginBiz) GoogleLogin(ctx context.Context, data authmdl.GoogleLoginDTO) (accessToken string, refreshToken string, err error) {
	user, err := biz.repo.GetUserByEmail(ctx, data.Email)
	if err != nil {

		passRand := common.GenSalt()
		salt := common.GenSalt()

		hash := hashser.Hash(passRand, salt)
		now := time.Now()
		newUser := authmdl.User{
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
		if err := biz.repo.CreateUser(ctx, &newUser); err != nil {
			biz.logger.Debug().Msgf("Biz errror 52")
			return "", "", common.ErrInternal(err)
		}
		accessToken = newUser.CreateAccessToken()
		refreshToken = newUser.CreateRefreshToken()

		go func() {
			defer common.AppRecover()
			if err := biz.repo.SaveRefreshToken(ctx, refreshToken, user.Id.Hex()); err != nil {
				panic(common.ErrSideEffectSaveRefreshToken(err))
			}
		}()

		return accessToken, refreshToken, nil
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
