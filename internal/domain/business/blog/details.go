package blogBusiness

import (
	"context"

	"nta-blog/internal/common"
	blogModel "nta-blog/internal/domain/model/blog"
	tagModel "nta-blog/internal/domain/model/tag"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type detailsBlogService interface {
	FindDetailsBlog(ctx context.Context, conditions map[string]interface{}) (*blogModel.Blog, error)
	ListBlog(ctx context.Context, pipeline bson.A) ([]blogModel.Blog, error)
	ListTag(ctx context.Context, filter map[string]interface{}) ([]tagModel.Tag, error)
}

type detailsBlogBiz struct {
	service detailsBlogService
}

func NewDetailsBlogBiz(sv detailsBlogService) *detailsBlogBiz {
	return &detailsBlogBiz{service: sv}
}

func (biz *detailsBlogBiz) FindDetailsBlog(ctx context.Context, blogId primitive.ObjectID, isForMetadata bool) (map[string]interface{}, error) {
	var result map[string]interface{}
	var tagDetails []map[string]interface{}
	blog, err := biz.service.FindDetailsBlog(ctx, map[string]interface{}{"_id": blogId})
	if err != nil {
		return nil, err
	}

	if isForMetadata {
		result = map[string]interface{}{
			"id":        blog.Id.Hex(),
			"title":     blog.Title,
			"summary":   blog.Summary,
			"thumbnail": blog.Thumbnail,
		}
	} else {
		var pipeline bson.A = bson.A{}
		pipeline = append(pipeline, bson.M{"$match": bson.M{"status": 1, "type": blog.Type, "_id": bson.M{"$ne": blogId}}})
		pipeline = append(pipeline, bson.M{"$limit": 5})

		pipeline = append(pipeline, bson.M{"$project": bson.M{
			"_id":       1,
			"title":     1,
			"thumbnail": 1,
		}})
		pipeline = append(pipeline, bson.M{"$sort": bson.M{"_id": -1}})
		queryResult, err := biz.service.ListBlog(ctx, pipeline)
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		relatedBlog := []map[string]interface{}{}
		if len(queryResult) != 0 {
			for _, blog := range queryResult {
				relatedBlog = append(relatedBlog, map[string]interface{}{
					"id":        blog.Id.Hex(),
					"title":     blog.Title,
					"thumbnail": blog.Thumbnail,
				})
			}
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
			"created_at": blog.CreatedAt,
			"updated_at": blog.UpdatedAt,
			"related":    relatedBlog,
		}
	}

	return result, nil
}
