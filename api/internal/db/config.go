package db

import (
	"app/internal/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConnection() *gorm.DB {
	dsn := "user:password@tcp(db:3306)/quizz-golang?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.User{})
	if err != nil {
		panic("Failed migration")
	}

	return db
}
