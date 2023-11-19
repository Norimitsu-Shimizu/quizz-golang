package diprovider

import (
	echoserver "app/internal"
	"app/internal/controller"
	"app/internal/db"
	adapter "app/internal/infrastructure"
	"app/internal/router"
	"app/internal/server"
	"app/internal/usecase"
	"context"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Providers = []fx.Option{
	EchoServerProviders,
	ControllerProviders,
	UsecaseProviders,
	RepositoryProviders,
	DBProviders,
	fx.Invoke(RegisterHooks),
	fx.Invoke(router.Router),
}

var EchoServerProviders = fx.Provide(
	server.NewServer,
	echoserver.StartServer,
)

var ControllerProviders = fx.Provide(
	controller.NewSignUpController,
	controller.NewSignInController,
)

var UsecaseProviders = fx.Provide(
	usecase.NewSignUpUsecase,
	usecase.NewSignInUsecase,
)

var RepositoryProviders = fx.Provide(
	adapter.NewUserGateway,
)

var DBProviders = fx.Provide(
	db.NewDatabaseConnection,
)

func RegisterHooks(lifecycle fx.Lifecycle, e *echo.Echo) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go e.Start(":8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
