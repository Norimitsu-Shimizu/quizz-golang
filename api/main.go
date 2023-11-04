package main

import (
	"app/internal/domain"
	"app/internal/router"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router.Router()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		db := dbInit()
		db.AutoMigrate(&domain.User{})
		return c.String(http.StatusOK, "Hello, world!")
	})
	e.Logger.Fatal(e.Start(":8080"))

	// db := dbInit()
	// db.AutoMigrate(&domain.User{})
}

func dbInit() *gorm.DB {
	dsn := "user:password@tcp(db:3306)/quizz-golang?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
