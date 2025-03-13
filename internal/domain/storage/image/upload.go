package imageStorage

import (
	"context"
	"encoding/json"
	"mime/multipart"

	imageModel "nta-blog/internal/domain/model/image"

	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func (s *store) UploadFile(ctx context.Context, file multipart.File) (*imageModel.UploadResFormCld, error) {
	var res imageModel.UploadResFormCld
	uploadParam, err := s.cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "/blog-image"})
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(uploadParam)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
