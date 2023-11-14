package adapter

import (
	"app/internal/domain"
	"app/internal/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserGateway struct {
	DB *gorm.DB
}

func NewUserGateway(db *gorm.DB) repository.UserRepository {
	return &UserGateway{
		DB: db,
	}
}

func (u *UserGateway) Create(c echo.Context, user *domain.User) error {
	if err := u.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
