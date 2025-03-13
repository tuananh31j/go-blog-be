package imageModel

import (
	"path/filepath"
	"strings"
)

const ImageCollection = "images"

type Image struct {
	SecureUrl string `bson:"secure_url" json:"secure_url"`
	PublicId  string `bson:"public_id" json:"public_id"`
	Width     uint16 `bson:"width" json:"width"`
	Height    uint16 `bson:"height" json:"height"`
}

func (i *Image) IsImage(filename string) bool {
	allowedExts := []string{".jpg", ".jpeg", ".png", ".gif"}
	ext := strings.ToLower(filepath.Ext(filename))
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			return true
		}
	}
	return false
}
