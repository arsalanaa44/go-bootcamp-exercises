package httpserver

import (
	"fmt"
	"gameapp/config"
	"gameapp/service/authservice"
	"gameapp/service/userservice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config  config.Config
	authSvc authservice.Service
	userSvc userservice.Service
}

func New(
	config config.Config,
	authSvc authservice.Service,
	userSvc userservice.Service,
) Server {

	return Server{
		config,
		authSvc,
		userSvc,
	}
}

func (s Server) Serve() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health-check", s.HealthCheck)
	userG := e.Group("/users")
	userG.POST("/register", s.userRegister)
	userG.POST("/login", s.userLogin)
	userG.POST("/profile", s.userProfile)
	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))

}
