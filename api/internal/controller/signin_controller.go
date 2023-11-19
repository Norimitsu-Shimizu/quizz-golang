package controller

import (
	"app/internal/parameter"
	"app/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SignInController struct {
	signInUsecase usecase.ISignInUsecase
}

func NewSignInController(signInUsecase usecase.ISignInUsecase) *SignInController {
	return &SignInController{
		signInUsecase: signInUsecase,
	}
}

func (s *SignInController) SignIn(c echo.Context) error {
	body := c.Request().Body

	if body == nil {
		return c.String(http.StatusBadRequest, "bodyがありません")
	}

	var req *parameter.PostLoginRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := s.signInUsecase.Execute(c, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
