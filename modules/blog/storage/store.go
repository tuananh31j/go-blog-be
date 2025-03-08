package blogstrg

import "gorm.io/gorm"

type splStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *splStore {
	return &splStore{db: db}
}
