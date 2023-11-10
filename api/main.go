package main

import (
	diprovider "app/internal/di"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	// router.Router()
	fxApp := fx.New(diprovider.Providers...)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		// db := dbInit()
		// db.AutoMigrate(&domain.User{})
		return c.String(http.StatusOK, "Hello, world!")
	})
	e.Logger.Fatal(e.Start(":8080"))
	fxApp.Run()

	// db := dbInit()
	// db.AutoMigrate(&domain.User{})
}

// func dbInit() *gorm.DB {
// 	dsn := "user:password@tcp(db:3306)/quizz-golang?charset=utf8mb4&parseTime=true&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	return db
// }
