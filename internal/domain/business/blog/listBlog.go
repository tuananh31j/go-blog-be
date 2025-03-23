package blogBusiness

import (
	"context"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	blogModel "nta-blog/internal/domain/model/blog"

	"go.mongodb.org/mongo-driver/bson"
)

type ListBlogService interface {
	List(ctx context.Context, pipeline bson.A) ([]blogModel.Blog, error)
}

type listBlogBiz struct {
	service ListBlogService
}

func NewListBlogBiz(service ListBlogService) *listBlogBiz {
	return &listBlogBiz{service: service}
}

func (biz *listBlogBiz) ListBlog(ctx context.Context, blogType cnst.IBlogType) ([]map[string]interface{}, error) {
	var pipeline bson.A = bson.A{}
	pipeline = append(pipeline, bson.M{"$match": bson.M{"status": 1, "type": blogType}})
	pipeline = append(pipeline, bson.M{"$lookup": bson.M{
		"from":         "tags",
		"localField":   "tag_ids",
		"foreignField": "_id",
		"as":           "tags",
	}})

	pipeline = append(pipeline, bson.M{"$lookup": bson.M{
		"from":         "users",
		"localField":   "user_id",
		"foreignField": "_id",
		"as":           "user",
	}})
	pipeline = append(pipeline, bson.M{"$unwind": bson.M{
		"path":                       "$user",
		"preserveNullAndEmptyArrays": true,
	}})
	pipeline = append(pipeline, bson.M{"$project": bson.M{
		"title":     1,
		"summary":   1,
		"thumbnail": 1,
		"tags": bson.M{"$map": bson.M{
			"input": "$tags",
			"as":    "tag",
			"in": bson.M{
				"_id":  "$$tag._id",
				"name": "$$tag.name",
			},
		}},
		"created_at": 1,
		"user": bson.M{
			"name_fake": 1,
		},
	}})
	pipeline = append(pipeline, bson.M{"$sort": bson.M{"_id": -1}})
	queryResult, err := biz.service.List(ctx, pipeline)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	responseData := make([]map[string]interface{}, 0)
	for _, blog := range queryResult {
		blogMap := map[string]interface{}{
			"id":         blog.Id.Hex(),
			"title":      blog.Title,
			"summary":    blog.Summary,
			"thumbnail":  blog.Thumbnail,
			"tags":       blog.Tags,
			"created_at": blog.CreatedAt,
		}
		responseData = append(responseData, blogMap)
	}

	return responseData, nil
}
