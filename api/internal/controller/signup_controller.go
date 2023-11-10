package controller

import (
	"app/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SignUpController struct {
	signUpUsecase *usecase.SignUpUsecase
}

func NewSignUpController(signUpUsecase *usecase.SignUpUsecase) *SignUpController {
	return &SignUpController{
		signUpUsecase: signUpUsecase,
	}
}

// func (s *SignUpController) SignUp(c echo.Context) {

// }

func (s *SignUpController) SignUp(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}
