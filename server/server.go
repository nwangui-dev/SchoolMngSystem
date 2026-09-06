package server

import (
	"net/http"

	"github.com/nwangui-dev/SchoolMngSystem/configs"
	"github.com/nwangui-dev/SchoolMngSystem/controllers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo   *echo.Echo
	config *configs.Config
}

func NewServer(config *configs.Config) *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &CustomValidator{Validator: validator.New()}

	controllers.InitRouter(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "School Management System API Service is healthy.")
	})

	return &Server{
		echo:   e,
		config: config,
	}
}

func (s *Server) Start() {
	port := s.config.App.Port
	if port != "" && port[0] != ':' {
		port = ":" + port
	}
	s.echo.Logger.Fatal(s.echo.Start(port))
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}