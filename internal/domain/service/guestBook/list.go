package guestbookService

import (
	"context"

	guestbookModel "nta-blog/internal/domain/model/guestBook"

	"go.mongodb.org/mongo-driver/bson"
)

type ListGuestBookStore interface {
	List(ctx context.Context, pipeline bson.A) ([]guestbookModel.GuestBook, error)
}

type listGuestBookService struct {
	store ListGuestBookStore
}

func NewListGuestBookService(s ListGuestBookStore) *listGuestBookService {
	return &listGuestBookService{store: s}
}

func (sv *listGuestBookService) GetListMessage(ctx context.Context, pipeline bson.A) ([]guestbookModel.GuestBook, error) {
	return sv.store.List(ctx, pipeline)
}

func (sv *listGuestBookService) TotalDocs(ctx context.Context, pipeline bson.A) (uint32, error) {
	filteredPipeline := bson.A{}

	for _, stage := range pipeline {
		stageMap, ok := stage.(bson.M)
		if !ok {
			continue
		}

		// Kiểm tra nếu stage là $skip hoặc $limit thì bỏ qua
		if _, isSkip := stageMap["$skip"]; isSkip {
			continue
		}
		if _, isLimit := stageMap["$limit"]; isLimit {
			continue
		}

		// Thêm vào pipeline mới
		filteredPipeline = append(filteredPipeline, stage)
	}

	pipeline = filteredPipeline
	result, err := sv.store.List(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	totalDocs := len(result)
	return uint32(totalDocs), nil
}
