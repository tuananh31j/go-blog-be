package guestBookHttp

import (
	"math"
	"strconv"

	"nta-blog/internal/common"
	guestbookBusiness "nta-blog/internal/domain/business/guestBook"
	"nta-blog/internal/domain/service"
	guestBookStorage "nta-blog/internal/domain/storage/guestBook"
	userStorage "nta-blog/internal/domain/storage/user"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func ListMessage(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		logger := apctx.GetLogger()
		mongodb := apctx.GetMongoDB()
		rdb := apctx.GetRedis()
		paging := c.Query("page", "1")
		limitStr := c.Query("limit", "10")

		guestBookStore := guestBookStorage.NewStore(mongodb)
		userStore := userStorage.NewStore(mongodb, rdb)
		service := service.NewGuestBookService(guestBookStore, userStore)
		biz := guestbookBusiness.NewListGuestBookBiz(service)
		limitInt, err := strconv.Atoi(limitStr)
		limit := uint32(limitInt)
		if err != nil {
			limit = 10
		}
		pageInt, err := strconv.Atoi(paging)
		page := uint32(pageInt)
		if err != nil {
			page = 1
		}
		data, total, err := biz.GetListMessage(c.Context(), page, limit)
		if err != nil {
			logger.Err(err).Msg("Failed to get list message")
			panic(err)
		}
		totalPage := math.Ceil(float64(total) / float64(limit))

		return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(data, page, totalPage, total, map[string]interface{}{"limit": limit, "page": page}))
	}
}
