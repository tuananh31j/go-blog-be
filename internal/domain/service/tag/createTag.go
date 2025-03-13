package tagService

import (
	"context"

	tagModel "nta-blog/internal/domain/model/tag"
)

type CreateTagStore interface {
	GetOnceTag(ctx context.Context, conditions map[string]interface{}) (*tagModel.Tag, error)
	CreateTag(ctx context.Context, dto tagModel.TagDTO) error
}

type createTagService struct {
	store CreateTagStore
}

func NewCreateTagService(s CreateTagStore) *createTagService {
	return &createTagService{store: s}
}

func (sv *createTagService) CheckTagExists(ctx context.Context, tagName string) error {
	_, err := sv.store.GetOnceTag(ctx, map[string]interface{}{"name": tagName})
	return err
}

func (sv *createTagService) CreateNewTag(ctx context.Context, tag tagModel.TagDTO) error {
	return sv.store.CreateTag(ctx, tag)
}
