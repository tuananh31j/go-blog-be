package authHttp

import (
	"nta-blog/internal/common"
	authBusiness "nta-blog/internal/domain/business/auth"
	"nta-blog/internal/domain/service"
	tokenStorage "nta-blog/internal/domain/storage/token"
	userStorage "nta-blog/internal/domain/storage/user"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Confirm(actx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var payload struct {
			Name string `json:"name"`
		}
		logger := actx.GetLogger()

		if err := c.BodyParser(&payload); err != nil {
			logger.Error().Err(err).Msg("Failed to parse request body")
			panic(common.ErrBadRequest(err))
		}
		userId, ok := c.Locals("userId").(primitive.ObjectID)
		if !ok {
			logger.Debug().Msg("Failed to get userId from context")
			panic(common.ErrBadRequest(nil))
		}
		rdb := actx.GetRedis()
		mongodb := actx.GetMongoDB()
		userStore := userStorage.NewStore(mongodb, rdb)
		tokenStore := tokenStorage.NewStore(mongodb, rdb)
		confirmService := service.NewAuthService(userStore, tokenStore)
		biz := authBusiness.NewConfirmBiz(confirmService)

		refreshToken, userTiny, err := biz.Confirm(c.Context(), userId, payload.Name)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to confirm user")
			panic(err)
		}

		return c.Status(fiber.StatusCreated).JSON(common.SimpleSuccessResponse(map[string]interface{}{"refreshToken": refreshToken, "userTiny": userTiny}))
	}
}
