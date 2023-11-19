package router

import (
	"app/internal/controller"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(e *echo.Echo, sc *controller.SignUpController, si *controller.SignInController) {
	// e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	api := e.Group("/api/v1")

	api.POST("/signup", sc.SignUp)
	api.POST("/signin", si.SignIn)

}
