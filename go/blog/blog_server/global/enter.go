package global

import (
	"blog_server/conf"

	"gorm.io/gorm"
)

var (
	Config *conf.Config
	DB     *gorm.DB
)
