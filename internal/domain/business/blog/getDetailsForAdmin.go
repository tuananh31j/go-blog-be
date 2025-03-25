package blogBusiness

import (
	"context"

	"nta-blog/internal/common"
	blogModel "nta-blog/internal/domain/model/blog"
	tagModel "nta-blog/internal/domain/model/tag"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type getDetailAllService interface {
	FindDetailsBlog(ctx context.Context, conditions map[string]interface{}) (*blogModel.Blog, error)
	ListTag(ctx context.Context, filter map[string]interface{}) ([]tagModel.Tag, error)
}

type getDetailAllBiz struct {
	service getDetailAllService
}

func NewGetForAdminBiz(sv getDetailAllService) *getDetailAllBiz {
	return &getDetailAllBiz{service: sv}
}

func (biz *getDetailAllBiz) FindDetailsALlBlog(ctx context.Context, blogId primitive.ObjectID) (map[string]interface{}, error) {
	var result map[string]interface{}
	var tagDetails []map[string]interface{}
	blog, err := biz.service.FindDetailsBlog(ctx, map[string]interface{}{"_id": blogId})
	if err != nil {
		return nil, err
	}

	for _, tagId := range blog.TagIds {
		tag, err := biz.service.ListTag(ctx, map[string]interface{}{"_id": tagId})
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		resultTag := map[string]interface{}{
			"id":   tag[0].Id.Hex(),
			"name": tag[0].Name,
		}
		tagDetails = append(tagDetails, resultTag)

	}
	result = map[string]interface{}{
		"id":         blog.Id.Hex(),
		"title":      blog.Title,
		"summary":    blog.Summary,
		"content":    blog.Content,
		"thumbnail":  blog.Thumbnail,
		"tags":       tagDetails,
		"status":     blog.Status,
		"user_id":    blog.UserId,
		"type":       blog.Type,
		"created_at": blog.CreatedAt,
		"updated_at": blog.UpdatedAt,
	}

	return result, nil
}
