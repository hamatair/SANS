package mysql

import (
	"SANS/entity"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(entity.User{})
}
