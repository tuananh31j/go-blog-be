package authtransportfiber

import (
	"time"

	"nta-blog/common"
	"nta-blog/config"
	"nta-blog/libs/appctx"
	authBusiness "nta-blog/modules/auth/business"
	authmdl "nta-blog/modules/auth/model"
	authrepo "nta-blog/modules/auth/repo"
	authtore "nta-blog/modules/auth/storage"

	"github.com/gofiber/fiber/v2"
)

func Login(actx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var payload authmdl.LoginDTO
		if err := c.BodyParser(&payload); err != nil {
			panic(common.ErrBadRequest(err))
		}
		logger := actx.GetLogger()
		rdb := actx.GetRedis()
		mongodb := actx.GetMongoDB()
		store := authtore.NewStore(mongodb, rdb)
		repo := authrepo.NewLoginRepo(store)
		biz := authBusiness.NewLoginBiz(repo, logger)

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
