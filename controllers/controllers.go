package controllers

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/nwangui-dev/SchoolMngSystem/handlers"
)

func InitRouter(e *echo.Echo) {
	api := e.Group("/api/v1")

	// Public Routes
	api.POST("/login", handlers.Login)

	// Configure JWT Middleware
	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(handlers.JwtCustomClaims)
		},
		SigningKey: handlers.JwtSecret,
	}

	// Protected Routes (Requires valid Authorization: Bearer <token>)
	protected := api.Group("")
	protected.Use(echojwt.WithConfig(jwtConfig))

	protected.POST("/students", handlers.CreateStudent)
	protected.GET("/students", handlers.GetStudents)
	protected.GET("/students/:id", handlers.GetStudentByID)
	protected.PUT("/students/:id", handlers.UpdateStudent)
	protected.DELETE("/students/:id", handlers.DeleteStudent)
}