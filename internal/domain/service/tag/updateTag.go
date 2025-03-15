package tagService

import (
	"context"

	tagModel "nta-blog/internal/domain/model/tag"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateTagStore interface {
	Find(ctx context.Context, conditions map[string]interface{}) (*tagModel.Tag, error)
	Update(ctx context.Context, condition map[string]interface{}, dto map[string]interface{}) error
}

type updateTagService struct {
	store UpdateTagStore
}

func NewUpdateTagService(s UpdateTagStore) *updateTagService {
	return &updateTagService{store: s}
}

func (sv *updateTagService) CheckTagNameExists(ctx context.Context, tagName string) error {
	_, err := sv.store.Find(ctx, map[string]interface{}{"name": tagName})
	return err
}

func (sv *updateTagService) CheckTagIdExists(ctx context.Context, tagId primitive.ObjectID) (*tagModel.Tag, error) {
	result, err := sv.store.Find(ctx, map[string]interface{}{"_id": tagId})
	return result, err
}

func (sv *updateTagService) UpdateTag(ctx context.Context, tagId primitive.ObjectID, newName string) error {
	return sv.store.Update(ctx, map[string]interface{}{"_id": tagId}, map[string]interface{}{
		"name": newName,
	})
}
