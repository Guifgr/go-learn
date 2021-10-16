package database

import (
	"github.com/guifgr/go-learn/ApiFiberJWT/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	connection, err := gorm.Open(mysql.Open("root:root@/golearn"), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		panic("Could not automigrate")
	}
}
