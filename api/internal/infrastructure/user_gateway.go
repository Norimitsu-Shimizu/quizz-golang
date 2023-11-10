package adapter

import "app/internal/repository"

type UserGateway struct{}

func NewUserGateway() repository.UserRepository {
	return &UserGateway{}
}
