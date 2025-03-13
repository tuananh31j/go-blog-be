package imageBusiness

import (
	"context"
	"errors"
	"mime/multipart"

	"nta-blog/internal/common"
	imageModel "nta-blog/internal/domain/model/image"
)

type UploadImageService interface {
	UploadFile(ctx context.Context, file multipart.File) (*imageModel.UploadResFormCld, error)
	SaveImage(ctx context.Context, image *imageModel.Image) error
}

type uploadImageBiz struct {
	service UploadImageService
}

func NewUploadImageBiz(sv UploadImageService) *uploadImageBiz {
	return &uploadImageBiz{service: sv}
}

func (biz *uploadImageBiz) UploadImage(ctx context.Context, fileHeader *multipart.FileHeader) (*imageModel.Image, error) {
	const maxSize = 5 * 1024 * 1024 // 5MB
	var image imageModel.Image
	if fileHeader.Size > maxSize {
		return nil, common.NewErrorResponse(errors.New("Image size too large!"), "Your image too big!", "File too big")
	}
	if !image.IsImage(fileHeader.Filename) {
		return nil, common.ErrBadRequest(errors.New("This is not an image!"))
	}
	formFile, err := fileHeader.Open()
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	img, err := biz.service.UploadFile(ctx, formFile)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	imageData := imageModel.Image{
		SecureUrl: img.SecureURL,
		PublicId:  img.PublicID,
		Width:     img.Width,
		Height:    img.Height,
	}
	if err := biz.service.SaveImage(ctx, &imageData); err != nil {
		return nil, common.ErrInternal(err)
	}

	return &imageData, nil
}
