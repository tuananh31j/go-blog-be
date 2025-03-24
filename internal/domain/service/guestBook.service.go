package service

import (
	"context"

	repository "nta-blog/internal/domain/interface"
	guestbookModel "nta-blog/internal/domain/model/guestBook"
	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type guestBookService struct {
	guestBookStore repository.GuestRepository
	userStore      repository.UserRepository
}

func NewGuestBookService(gs repository.GuestRepository, us repository.UserRepository) *guestBookService {
	return &guestBookService{guestBookStore: gs, userStore: us}
}

func (sv *guestBookService) UpdateMessage(ctx context.Context, guestBookId primitive.ObjectID, updateField map[string]interface{}) error {
	return sv.guestBookStore.Update(ctx, guestBookId, updateField)
}

func (sv *guestBookService) FindMessage(ctx context.Context, condition map[string]interface{}) ([]*guestbookModel.GuestBook, error) {
	return sv.guestBookStore.Find(ctx, condition)
}

func (sv *guestBookService) FindOneUser(ctx context.Context, userId primitive.ObjectID) (*userModel.User, error) {
	return sv.userStore.FindOneUser(ctx, map[string]interface{}{"_id": userId})
}

func (sv *guestBookService) CreateMessage(ctx context.Context, dto *guestbookModel.GuestBook) error {
	return sv.guestBookStore.Create(ctx, dto)
}

func (sv *guestBookService) GetListMessage(ctx context.Context, pipeline bson.A) ([]guestbookModel.GuestBook, error) {
	return sv.guestBookStore.List(ctx, pipeline)
}

func (sv *guestBookService) TotalDocs(ctx context.Context, pipeline bson.A) (uint32, error) {
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
	result, err := sv.guestBookStore.List(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	totalDocs := len(result)
	return uint32(totalDocs), nil
}
