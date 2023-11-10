package diprovider

import (
	"app/internal/controller"
	"app/internal/db"
	adapter "app/internal/infrastructure"
	"app/internal/usecase"

	"go.uber.org/fx"
)

var Providers = []fx.Option{
	ControllerProviders,
	UsecaseProviders,
	RepositoryProviders,
	DBProviders,
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

var DBProviders = fx.Provide(
	db.NewDatabaseConnection,
)
