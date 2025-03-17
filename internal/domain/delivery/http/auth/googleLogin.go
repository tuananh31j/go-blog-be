package authHttp

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"nta-blog/internal/common"
	authBusiness "nta-blog/internal/domain/business/auth"
	userModel "nta-blog/internal/domain/model/user"
	authService "nta-blog/internal/domain/service/auth"
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
			logger.Debug().Msg("Lấy token thất bại")
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange code for token")
		}

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			return c.SendString("User Data Fetch Failed")
		}

		userData, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Debug().Msg("Lỗi đọc dữ liệu")
			return err
		}

		if err := json.Unmarshal(userData, &result); err != nil {
			logger.Debug().Msg("Lỗi unmarshal")
			return err
		}

		userStore := userStorage.NewStore(mongodb, rdb)
		tokenStore := tokenStorage.NewStore(mongodb, rdb)
		loginService := authService.NewLoginService(userStore, tokenStore)
		biz := authBusiness.NewGoogleLoginBiz(loginService, logger)

		accessToken, refreshToken, err := biz.GoogleLogin(c.Context(), &result)
		if err != nil {
			logger.Debug().Msg("Lỗi Google Login")
			return err
		}

		c.Cookie(&fiber.Cookie{
			Name:     "accessToken",
			Value:    accessToken,
			Path:     "/",
			HTTPOnly: true,
			Secure:   config.Env.AppENV == "production",
			SameSite: "None",
			Domain:   config.Env.AllowOrigin,
			Expires:  time.Now().Add(time.Minute * 15),
			MaxAge:   900,
		})

		c.Cookie(&fiber.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			Path:     "/",
			HTTPOnly: true,
			Secure:   config.Env.AppENV == "production",
			SameSite: "None",
			Domain:   config.Env.AllowOrigin,
			Expires:  time.Now().Add(24 * time.Hour * 7),
			MaxAge:   604800,
		})

		return c.Redirect(config.Env.NextJSRedirectOauth + accessToken + "ntadtt31012212" + common.GenSalt() + "?name=" + result.Name)
	}
}
