package echoserver

import (
	"github.com/labstack/echo/v4"
)

type EchoServer struct {
	Echo *echo.Echo
}

func StartServer(e *echo.Echo) *EchoServer {
	return &EchoServer{Echo: e}
}

func (s *EchoServer) Start(address string) error {
	return s.Echo.Start(address)
}
