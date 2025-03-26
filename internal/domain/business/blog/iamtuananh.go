package blogBusiness

import (
	"context"

	cnst "nta-blog/internal/constant"
	blogModel "nta-blog/internal/domain/model/blog"
)

type iamtuananhService interface {
	FindDetailsBlog(ctx context.Context, conditions map[string]interface{}) (*blogModel.Blog, error)
}

type iamtuananhBiz struct {
	service iamtuananhService
}

func NewIamtuananhBiz(sv iamtuananhService) *iamtuananhBiz {
	return &iamtuananhBiz{service: sv}
}

func (biz *iamtuananhBiz) Iamtuananh(ctx context.Context) (map[string]interface{}, error) {
	var result map[string]interface{}
	blog, err := biz.service.FindDetailsBlog(ctx, map[string]interface{}{"type": cnst.BlogTypeConstant.Me})
	if err != nil {
		return nil, err
	}

	result = map[string]interface{}{
		"id":         blog.Id.Hex(),
		"title":      blog.Title,
		"summary":    blog.Summary,
		"content":    blog.Content,
		"thumbnail":  blog.Thumbnail,
		"created_at": blog.CreatedAt,
		"updated_at": blog.UpdatedAt,
	}

	return result, nil
}
