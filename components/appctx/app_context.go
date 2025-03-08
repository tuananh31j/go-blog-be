package appctx

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMyDBConnection() *gorm.DB
}

type appctx struct {
	db     *gorm.DB
	logger *zerolog.Logger
}

func NewAppContext(db *gorm.DB, lg *zerolog.Logger) *appctx {
	return &appctx{db: db, logger: lg}
}

func (actx *appctx) GetMyDBConnection() *gorm.DB {
	return actx.db
}

func (actx *appctx) GetLogger() *zerolog.Logger {
	return actx.logger
}
