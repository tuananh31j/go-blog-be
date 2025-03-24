package authHttp

import (
	"errors"

	"nta-blog/internal/common"
	authBusiness "nta-blog/internal/domain/business/auth"
	"nta-blog/internal/domain/service"
	tokenStorage "nta-blog/internal/domain/storage/token"
	userStorage "nta-blog/internal/domain/storage/user"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func RefreshToken(actx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		rdb := actx.GetRedis()
		logger := actx.GetLogger()
		tokenStore := tokenStorage.NewStore(actx.GetMongoDB(), rdb)
		userStore := userStorage.NewStore(actx.GetMongoDB(), rdb)
		service := service.NewAuthService(userStore, tokenStore)
		biz := authBusiness.NewRefreshTokenBiz(service)
		refreshToken := c.Cookies("refreshToken")
		if refreshToken == "" {
			logger.Debug().Msg("refreshToken not found")
			panic(common.ErrBadRequest(errors.New("refreshToken not found")))
		}

		accessToken, err := biz.RefreshToken(c.Context(), refreshToken)
		if err != nil {
			logger.Debug().Msg(err.Error())
			panic(err)
		}

		c.Status(fiber.StatusCreated)

		return c.JSON(common.SimpleSuccessResponse(map[string]string{"token": accessToken}))
	}
}
