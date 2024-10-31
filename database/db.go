// database/db.go
package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"dating-app/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "user:password@tcp(127.0.0.1:3306)/datingapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	db.AutoMigrate(&models.User{}, &models.SwipeHistory{})
}
