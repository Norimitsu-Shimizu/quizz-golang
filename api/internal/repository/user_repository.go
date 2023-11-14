package repository

import (
	"app/internal/domain"

	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	Create(c echo.Context, user *domain.User) error
}
