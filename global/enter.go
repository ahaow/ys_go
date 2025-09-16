package global

import (
	"ys_go/config"
	"ys_go/initialize/logger"

	"gorm.io/gorm"
)

var (
	Config *config.Config
	Log    *logger.Logger
	DB     *gorm.DB
)
