package middlewares

import (
	"nta-blog/common"
	"nta-blog/components/appctx"

	"github.com/gofiber/fiber/v2"
)

func Recover(actx appctx.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error { // Thêm error để phù hợp với fiber.Handler
		defer func() { // Thêm () để gọi hàm vô danh
			if err := recover(); err != nil {
				c.Request().Header.Set("Content-Type", "application/json") // Sửa cú pháp Header

				if appErr, ok := err.(*common.AppError); ok {
					c.Status(appErr.StatusCode).JSON(appErr)
					// Gọi lại panic để framework bắt lỗi
					panic(appErr)
				}
				appErr := common.ErrInternal(err.(error))
				c.Status(appErr.StatusCode).JSON(appErr)
				panic(appErr)
			}
		}()

		return c.Next() // Trả về error từ c.Next()
	}
}
