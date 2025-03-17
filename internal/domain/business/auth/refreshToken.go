package authBusiness

import (
	"context"
	"strconv"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	userModel "nta-blog/internal/domain/model/user"
	"nta-blog/internal/infrastructure/config"
	"nta-blog/internal/lib/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefreshTokenService interface {
	RemoveRefreshToken(ctx context.Context, token string) error
	CheckRefreshTokenExists(ctx context.Context, userId string) (string, error)
}

type refreshTokenBiz struct {
	service RefreshTokenService
}

func NewRefreshTokenBiz(sv RefreshTokenService) *refreshTokenBiz {
	return &refreshTokenBiz{service: sv}
}

func (biz *refreshTokenBiz) RefreshToken(ctx context.Context, token string) (accessToken string, err error) {
	var user userModel.User

	payload, errVeri := user.VerifyToken(token, config.Env.SecretRefreshKey)
	if errVeri != nil {
		return "", common.ErrBadRequest(err)
	}

	currentToken, errBiz := biz.service.CheckRefreshTokenExists(ctx, payload.Id)
	if errBiz != nil {
		logger.ZeroLog.Debug().Err(err).Msgf("CheckRefreshTokenExists: %v", payload.Id)
		return "", common.ErrBadRequest(err)
	}
	if token != currentToken {
		return "", common.ErrBadRequest(nil)
	}

	objectId, errParse := primitive.ObjectIDFromHex(payload.Id)
	if errParse != nil {
		return "", common.ErrBadRequest(err)
	}

	user.Id = objectId
	role, err := strconv.Atoi(payload.Role)
	if err != nil {
		return "", common.ErrBadRequest(err)
	}
	user.Role = cnst.TRoleAccount(role)

	accessToken = user.CreateAccessToken()

	return accessToken, nil
}
