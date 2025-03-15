package authBusiness

import (
	"context"
	"strconv"
	"time"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	userModel "nta-blog/internal/domain/model/user"
	"nta-blog/internal/infrastructure/config"
	"nta-blog/internal/lib/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefreshTokenService interface {
	SaveRefreshToken(ctx context.Context, token string) error
	RemoveRefreshToken(ctx context.Context, token string) error
	CheckRefeshTokenExists(ctx context.Context, token string) (string, error)
}

type refreshTokenBiz struct {
	service RefreshTokenService
}

func NewRefreshTokenBiz(sv RefreshTokenService) *refreshTokenBiz {
	return &refreshTokenBiz{service: sv}
}

func (biz *refreshTokenBiz) SaveRefreshToken(ctx context.Context, token string) (accessToken string, refreshToken string, err error) {
	var user userModel.User

	oldToken, err := biz.service.CheckRefeshTokenExists(ctx, token)
	if token != oldToken {
		return "", "", common.ErrBadRequest(nil)
	}
	if err != nil {
		return "", "", common.ErrBadRequest(err)
	}
	payload, err := user.VerifyToken(token, config.Env.SecretRefreshKey)
	if err != nil {
		return "", "", common.ErrBadRequest(err)
	}
	objectId, err := primitive.ObjectIDFromHex(payload.Id)
	if err != nil {
		return "", "", common.ErrBadRequest(err)
	}

	user.Id = objectId
	role, err := strconv.Atoi(payload.Role)
	user.Role = cnst.TRoleAccount(role)

	accessToken = user.CreateAccessToken()
	refreshToken = user.CreateRefreshToken()
	if err := biz.service.SaveRefreshToken(ctx, refreshToken); err != nil {
		return "", "", common.ErrInternal(err)
	}
	go func() {
		defer common.AppRecover()
		retries := 3
		for i := 0; i < retries; i++ {
			if err := biz.service.RemoveRefreshToken(ctx, token); err == nil {
				break
			}
			if i == retries-1 {
				time.Sleep(time.Duration(i) * time.Second)
				logger.ZeroLog.Error().Err(err).Msg("Failed to remove refresh token after retries")
			}
		}
	}()
	return accessToken, refreshToken, nil
}
