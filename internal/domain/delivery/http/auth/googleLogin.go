package authHttp

import (
	"encoding/json"
	"io"
	"net/http"
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

func GoogleLogin(actx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var result userModel.GoogleLoginDTO
		rdb := actx.GetRedis()
		mongodb := actx.GetMongoDB()
		logger := actx.GetLogger()

		googleCon := config.GuConfig.GoogleLoginConfig

		code := c.Query("code")

		token, err := googleCon.Exchange(c.Context(), code)
		if err != nil {
			logger.Debug().Msg("Lấy token")

			return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange code for token")
		}

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			return c.SendString("User Data Fetch Failed")
		}

		userData, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Debug().Msg("Lỗi json")
			return err
		}

		if err := json.Unmarshal(userData, &result); err != nil {
			logger.Debug().Msg("Lấy unmarshal")

			return err
		}

		userStore := userStorage.NewStore(mongodb, rdb)
		tokenStore := tokenStorage.NewStore(mongodb, rdb)
		loginSevice := loginService.NewLoginService(userStore, tokenStore)
		biz := authBusiness.NewGoogleLoginBiz(loginSevice, logger)

		var accessToken, refreshToken string
		accessToken, refreshToken, err = biz.GoogleLogin(c.Context(), result)
		if err != nil {
			logger.Debug().Msg("")

			return err
		}

		c.Cookie(&fiber.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			Expires:  time.Now().Add(24 * time.Hour * 7),
			HTTPOnly: true,
			Secure:   config.Env.AppENV != "development",
		})

		if err := c.Status(fiber.StatusOK).JSON(common.SimpleSuccessResponse(fiber.Map{
			"token": accessToken,
		})); err != nil {
			return err
		}
		return nil
	}
}
