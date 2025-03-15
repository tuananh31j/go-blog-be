package authHttp

import (
	"errors"

	"nta-blog/internal/common"
	authBusiness "nta-blog/internal/domain/business/auth"
	authService "nta-blog/internal/domain/service/auth"
	tokenStorage "nta-blog/internal/domain/storage/token"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func RefreshToken(actx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		rdb := actx.GetRedis()
		logger := actx.GetLogger()
		tokenStore := tokenStorage.NewStore(actx.GetMongoDB(), rdb)
		service := authService.NewRefreshTokenService(tokenStore)
		biz := authBusiness.NewRefreshTokenBiz(service)
		oldRefreshToken := c.Cookies("refreshToken")
		if oldRefreshToken == "" {
			logger.Debug().Msg("refreshToken not found")
			panic(common.ErrBadRequest(errors.New("refreshToken not found")))
		}

		accessToken, refreshToken, err := biz.SaveRefreshToken(c.Context(), oldRefreshToken)
		if err != nil {
			logger.Debug().Msg(err.Error())
			panic(err)
		}

		c.Status(fiber.StatusCreated)
		if err := c.JSON(common.SimpleSuccessResponse(map[string]string{"token": accessToken, "refreshToken": refreshToken})); err != nil {
			logger.Debug().Msg(err.Error())
		}
		return nil
	}
}
