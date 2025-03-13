package tagService

import (
	"context"

	tagModel "nta-blog/internal/domain/model/tag"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DetailsTagStore interface {
	GetOnceTag(ctx context.Context, conditions map[string]interface{}) (*tagModel.Tag, error)
}

type detailsTagService struct {
	store DetailsTagStore
}

func NewDetailsTagService(s DetailsTagStore) *detailsTagService {
	return &detailsTagService{store: s}
}

func (sv *detailsTagService) GetDetailsTag(ctx context.Context, tagId primitive.ObjectID) error {
	_, err := sv.store.GetOnceTag(ctx, map[string]interface{}{"_id": tagId})
	return err
}
