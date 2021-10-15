package migrations

import (
	"github.com/guifgr/GO-LEARN/ApiRestDatabase/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
