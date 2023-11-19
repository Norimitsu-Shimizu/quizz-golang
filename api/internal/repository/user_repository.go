package repository

import (
	"app/internal/domain"

	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	Create(c echo.Context, user *domain.User) error
	GetByEmail(c echo.Context, email string) (*domain.User, error)
}
