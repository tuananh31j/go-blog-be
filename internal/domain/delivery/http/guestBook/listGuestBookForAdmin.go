package guestBookHttp

import (
	"strconv"

	"nta-blog/internal/common"
	guestbookBusiness "nta-blog/internal/domain/business/guestBook"
	"nta-blog/internal/domain/service"
	guestBookStorage "nta-blog/internal/domain/storage/guestBook"
	userStorage "nta-blog/internal/domain/storage/user"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func ListGuestBookForAdmin(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		logger := apctx.GetLogger()
		mongodb := apctx.GetMongoDB()
		rdb := apctx.GetRedis()
		paging := c.Query("page")
		limit := c.Query("limit")
		status := c.Query("status")
		logger.Debug().Msgf("paging: %s, limit: %s, status: %s", paging, limit, status)

		guestBookStore := guestBookStorage.NewStore(mongodb)
		userStore := userStorage.NewStore(mongodb, rdb)
		service := service.NewGuestBookService(guestBookStore, userStore)
		biz := guestbookBusiness.NewListGuestBookForAdminBiz(service)

		limitInt, err := strconv.Atoi(limit)
		length := uint32(limitInt)
		if err != nil {
			length = 10
		}
		pageInt, err := strconv.Atoi(paging)
		page := uint32(pageInt)
		if err != nil {
			page = 1
		}
		data, total, err := biz.GetListMessageForAdmin(c.Context(), page, length, status)
		if err != nil {
			logger.Err(err).Msg("Failed to get list message")
			panic(err)
		}
		totalPage := total / length

		return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(data, page, totalPage, total, map[string]interface{}{"limit": length, "status": status}))
	}
}
