package blogtspt

//	STRANSPORT LÀ NƠI HIỂU HẾT MỌI TẦNG
import (
	"nta-blog/components/appctx"
	"nta-blog/modules/blog/business"
	blogmdl "nta-blog/modules/blog/model"
	blogstrg "nta-blog/modules/blog/storage"

	"github.com/gofiber/fiber/v2"
)

func CreateBlog(actx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		store := blogstrg.NewSQLStore(actx.GetMyDBConnection())
		biz := business.NewCreateBlog(store)
		var blogData blogmdl.CreateBlog
		if err := c.BodyParser(&blogData); err != nil {
			return err
		}
		if err := biz.CreateBlog(c.Context(), &blogData); err != nil {
			return err
		}
		return nil
	}
}
