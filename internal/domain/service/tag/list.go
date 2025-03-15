package tagService

import (
	"context"

	tagModel "nta-blog/internal/domain/model/tag"
)

type ListTagStore interface {
	List(ctx context.Context, conditions map[string]interface{}) ([]tagModel.Tag, error)
}

type listTagService struct {
	tagStore ListTagStore
}

func NewListTagService(ts ListTagStore) *listTagService {
	return &listTagService{tagStore: ts}
}

func (s *listTagService) GetAllTag(ctx context.Context) ([]tagModel.Tag, error) {
	return s.tagStore.List(ctx, map[string]interface{}{})
}
