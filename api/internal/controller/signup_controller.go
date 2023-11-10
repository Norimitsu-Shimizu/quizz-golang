package controller

import (
	"app/internal/usecase"

	"github.com/labstack/echo/v4"
)

type signUpController struct {
	signUpUsecase usecase.SignUpUsecase
}

func NewSignUpController(signUpUsecase usecase.SignUpUsecase) *signUpController {
	return &signUpController{
		signUpUsecase: signUpUsecase,
	}
}

func (s signUpController) SignUp(c echo.Context) {

}
