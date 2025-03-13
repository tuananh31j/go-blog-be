package db

import "github.com/cloudinary/cloudinary-go"

func NewCld(name, key, secret string) (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(name, key, secret)
	return cld, err
}
