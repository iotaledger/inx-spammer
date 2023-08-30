package spammer

import (
	"github.com/labstack/echo/v4"
)

const (
	APIRoute = "/api/spammer/v1"
)

type Server struct {
	Spammer *Spammer
}

func NewServer(spammer *Spammer, echo *echo.Echo) *Server {
	s := &Server{
		Spammer: spammer,
	}
	s.configureRoutes(echo.Group(APIRoute))

	return s
}
