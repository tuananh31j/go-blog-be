package tagService

import (
	"context"

	tagModel "nta-blog/internal/domain/model/tag"
)

type CreateTagStore interface {
	Find(ctx context.Context, conditions map[string]interface{}) (*tagModel.Tag, error)
	Create(ctx context.Context, dto *tagModel.Tag) error
}

type createTagService struct {
	store CreateTagStore
}

func NewCreateTagService(s CreateTagStore) *createTagService {
	return &createTagService{store: s}
}

func (sv *createTagService) CheckTagExists(ctx context.Context, tagName string) error {
	_, err := sv.store.Find(ctx, map[string]interface{}{"name": tagName})
	return err
}

func (sv *createTagService) CreateNewTag(ctx context.Context, tag *tagModel.Tag) error {
	return sv.store.Create(ctx, tag)
}
