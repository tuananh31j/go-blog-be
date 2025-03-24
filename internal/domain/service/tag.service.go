package service

import (
	"context"

	repository "nta-blog/internal/domain/interface"
	tagModel "nta-blog/internal/domain/model/tag"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type tagService struct {
	tagStore repository.TagRepository
}

func NewTagService(ts repository.TagRepository) *tagService {
	return &tagService{tagStore: ts}
}

func (sv *tagService) CheckTagNameExists(ctx context.Context, tagName string) error {
	_, err := sv.tagStore.Find(ctx, map[string]interface{}{"name": tagName})
	return err
}

func (sv *tagService) CheckTagIdExists(ctx context.Context, tagId primitive.ObjectID) (*tagModel.Tag, error) {
	result, err := sv.tagStore.Find(ctx, map[string]interface{}{"_id": tagId})
	return result, err
}

func (sv *tagService) UpdateTag(ctx context.Context, tagId primitive.ObjectID, newName string) error {
	return sv.tagStore.Update(ctx, map[string]interface{}{"_id": tagId}, map[string]interface{}{
		"name": newName,
	})
}

func (s *tagService) GetAllTag(ctx context.Context) ([]tagModel.Tag, error) {
	return s.tagStore.List(ctx, map[string]interface{}{})
}

func (sv *tagService) GetDetailsTag(ctx context.Context, tagId primitive.ObjectID) error {
	_, err := sv.tagStore.Find(ctx, map[string]interface{}{"_id": tagId})
	return err
}

func (sv *tagService) CheckTagExists(ctx context.Context, tagName string) error {
	_, err := sv.tagStore.Find(ctx, map[string]interface{}{"name": tagName})
	return err
}

func (sv *tagService) CreateNewTag(ctx context.Context, tag *tagModel.Tag) error {
	return sv.tagStore.Create(ctx, tag)
}
