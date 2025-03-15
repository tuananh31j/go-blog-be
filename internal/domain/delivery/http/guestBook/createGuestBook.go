package guestBookHttp

import (
	"nta-blog/internal/common"
	guestbookBusiness "nta-blog/internal/domain/business/guestBook"
	guestbookService "nta-blog/internal/domain/service/guestBook"
	guestBookStorage "nta-blog/internal/domain/storage/guestBook"
	userStorage "nta-blog/internal/domain/storage/user"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateMessageRequest struct {
	Message string `json:"message"`
}

func CreateMessage(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		mongoDB := apctx.GetMongoDB()
		logger := apctx.GetLogger()
		rdb := apctx.GetRedis()
		var payload CreateMessageRequest
		if err := c.BodyParser(&payload); err != nil {
			logger.Debug().Err(err).Msg("Payload is not valid!")
			panic(err)
		}
		logger.Debug().Interface("payload", c.Locals("userId")).Msg("Payload")
		userId, ok := c.Locals("userId").(primitive.ObjectID)
		if !ok {
			logger.Debug().Msg("Failed to get userId from context")
			panic(common.ErrBadRequest(nil))
		}

		guestBookStore := guestBookStorage.NewStore(mongoDB)
		userStore := userStorage.NewStore(mongoDB, rdb)
		service := guestbookService.NewCreateGuestBookService(guestBookStore, userStore)
		biz := guestbookBusiness.NewCreateGuestBookBiz(service, logger)
		err := biz.CreateMessage(c.Context(), payload.Message, userId)
		if err != nil {
			logger.Debug().Err(err).Msg("Failed to create message")
			panic(err)
		}

		return c.Status(fiber.StatusCreated).JSON(common.SimpleSuccessResponse("Message created successfully"))
	}
}
