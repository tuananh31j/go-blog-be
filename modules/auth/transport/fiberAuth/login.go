package authtransportfiber

import (
	"time"

	"nta-blog/common"
	"nta-blog/config"
	"nta-blog/libs/appctx"
	"nta-blog/libs/hashser"
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
		store := authtore.NewStore(actx.GetMongoDB())
		repo := authrepo.NewLoginRepo(store)
		md5Hash := hashser.NewMd5Hash(payload.Password)
		biz := authBusiness.NewLoginBiz(repo, md5Hash)

		accessToken, refreshToken, err := biz.Login(c.Context(), payload)
		if err != nil {
			panic(err)
		}
		c.Cookie(&fiber.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
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
