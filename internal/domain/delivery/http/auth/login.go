package authHttp

import (
	"time"

	"nta-blog/internal/common"
	authBusiness "nta-blog/internal/domain/business/auth"
	userModel "nta-blog/internal/domain/model/user"
	authService "nta-blog/internal/domain/service/auth"
	tokenStorage "nta-blog/internal/domain/storage/token"
	userStorage "nta-blog/internal/domain/storage/user"
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
		loginSevice := authService.NewLoginService(userStore, tokenStore)
		biz := authBusiness.NewLoginBiz(loginSevice, logger)

		accessToken, refreshToken, userTiny, err := biz.Login(c.Context(), payload)
		if err != nil {
			panic(err)
		}

		c.Status(fiber.StatusCreated)

		userTinyMap := map[string]interface{}{
			"id":    userTiny.Id,
			"name":  userTiny.NameFake,
			"email": userTiny.Email,
			"role":  userTiny.Role,
			"avt":   userTiny.Avt,
		}
		c.Cookie(&fiber.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			Path:     "/",
			Domain:   "localhost",
			Expires:  time.Now().Add(24 * time.Hour * 7),
			HTTPOnly: true,
			Secure:   false,
			SameSite: "None",
		})
		return c.JSON(common.SimpleSuccessResponse(map[string]interface{}{"token": accessToken, "refreshToken": refreshToken, "userTiny": userTinyMap}))
	}
}
