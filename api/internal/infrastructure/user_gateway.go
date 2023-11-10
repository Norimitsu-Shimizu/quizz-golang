package adapter

import "app/internal/repository"

type userGateway struct{}

func NewUserGateway() repository.UserRepository {
	return &userGateway{}
}
