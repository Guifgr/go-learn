package migrations

import (
	"github.com/guifgr/GO-LEARN/ApiRestDatabase/models"
	"gorm.io/gorm"
	"log"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(models.Book{})
	if err != nil {
		log.Fatal("Migration Failed")
		return
	}
}
