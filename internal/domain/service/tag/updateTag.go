package tagService

import (
	"context"

	tagModel "nta-blog/internal/domain/model/tag"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateTagStore interface {
	GetOnceTag(ctx context.Context, conditions map[string]interface{}) (*tagModel.Tag, error)
	UpdateTag(ctx context.Context, tagId primitive.ObjectID, dto tagModel.TagDTO) error
}

type updateTagService struct {
	store UpdateTagStore
}

func NewUpdateTagService(s UpdateTagStore) *updateTagService {
	return &updateTagService{store: s}
}

func (sv *updateTagService) CheckTagNameExists(ctx context.Context, tagName string) error {
	_, err := sv.store.GetOnceTag(ctx, map[string]interface{}{"name": tagName})
	return err
}

func (sv *updateTagService) CheckTagIdExists(ctx context.Context, tagId primitive.ObjectID) (*tagModel.Tag, error) {
	result, err := sv.store.GetOnceTag(ctx, map[string]interface{}{"_id": tagId})
	return result, err
}

func (sv *updateTagService) UpdateTag(ctx context.Context, tagId primitive.ObjectID, tag tagModel.TagDTO) error {
	return sv.store.UpdateTag(ctx, tagId, tag)
}
