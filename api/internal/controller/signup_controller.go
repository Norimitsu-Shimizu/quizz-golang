package controller

import (
	"app/internal/parameter"
	"app/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SignUpController struct {
	signUpUsecase usecase.ISignUpUsecase
}

func NewSignUpController(signUpUsecase usecase.ISignUpUsecase) *SignUpController {
	return &SignUpController{
		signUpUsecase: signUpUsecase,
	}
}

func (s *SignUpController) SignUp(c echo.Context) error {
	body := c.Request().Body

	if body == nil {
		return c.String(http.StatusBadRequest, "bodyがありません")
	}

	var req *parameter.PostUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	s.signUpUsecase.Execute(c, req)

	// return c.String(http.StatusOK, "Hello, world!")
	return c.NoContent(http.StatusCreated)
}
