package tagModel

import "errors"

type TagDTO struct {
	Name string `bson:"name" json:"name"`
}

func (dto *TagDTO) ValidateNameNotEmpty() error {
	if dto.Name == "" {
		return errors.New("Tên thẻ không được để trống!")
	}
	return nil
}
