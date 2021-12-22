package global

import (
	"chuckie_gin/config"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Settings config.ServerConfig
	Lg *zap.Logger
	Trans ut.Translator
	DB *gorm.DB
)