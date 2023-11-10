package diprovider

import (
	"app/internal/controller"
	"app/internal/infrastructure"
	"app/internal/usecase"

	"go.uber.org/fx"
)

var Providers = []fx.Option{
	ControllerProviders,
	UsecaseProviders,
	RepositoryProviders,
}

var ControllerProviders = fx.Provide(
	controller.NewSignUpController,
)

var UsecaseProviders = fx.Provide(
	usecase.NewSignUpUsecase,
)

var RepositoryProviders = fx.Provide(
	adapter.NewUserGateway,
)
