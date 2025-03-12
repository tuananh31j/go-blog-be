package authHttp

import (
	"time"

	"nta-blog/internal/common"
	authBusiness "nta-blog/internal/domain/business/auth"
	userModel "nta-blog/internal/domain/model/user"
	loginService "nta-blog/internal/domain/service"
	tokenStorage "nta-blog/internal/domain/storage/token"
	userStorage "nta-blog/internal/domain/storage/user"
	"nta-blog/internal/infrastructure/config"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func Login(actx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var payload userModel.LoginDTO
		if err := c.BodyParser(&payload); err != nil {
			panic(common.ErrBadRequest(err))
		}
		logger := actx.GetLogger()
		rdb := actx.GetRedis()
		mongodb := actx.GetMongoDB()
		userStore := userStorage.NewStore(mongodb, rdb)
		tokenStore := tokenStorage.NewStore(mongodb, rdb)
		loginSevice := loginService.NewLoginService(userStore, tokenStore)
		biz := authBusiness.NewLoginBiz(loginSevice, logger)

		accessToken, refreshToken, err := biz.Login(c.Context(), payload)
		if err != nil {
			panic(err)
		}
		c.Cookie(&fiber.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			Path:     "/",
			Domain:   "localhost",
			Expires:  time.Now().Add(24 * time.Hour * 7),
			HTTPOnly: true,
			Secure:   config.Env.AppENV != "development",
		})

		c.Status(fiber.StatusCreated)
		if err := c.JSON(common.SimpleSuccessResponse(map[string]string{"token": accessToken})); err != nil {
			logger.Debug().Msg(err.Error())
		}
		return nil
	}
}
