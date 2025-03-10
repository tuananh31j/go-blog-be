package middleware

import (
	"fmt"

	"nta-blog/common"
	"nta-blog/libs/appctx"

	"github.com/gofiber/fiber/v2"
)

func Recover(actx appctx.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				logger := actx.GetLogger()
				logger.Debug().Msg(fmt.Sprintf("Recover>>>>>>> %v", err))

				c.Set("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.Status(appErr.StatusCode)
					errFiber := c.JSON(appErr)
					if errFiber != nil {
						logger.Debug().Msg("Faild to send json response!")
					}
				} else {
					appErr := common.ErrInternal(err.(error))
					c.Status(appErr.StatusCode)
					errFiber := c.JSON(appErr)
					if errFiber != nil {
						logger.Debug().Msg("Faild to send json response!")
					}
				}
			}
		}()
		return c.Next()
	}
}
