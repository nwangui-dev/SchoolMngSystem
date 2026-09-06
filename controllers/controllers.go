package controllers

import (
	"github.com/nwangui-dev/SchoolMngSystem/handlers"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	api := e.Group("/api/v1")

	// Student endpoints
	api.POST("/students", handlers.CreateStudent)
	api.GET("/students", handlers.GetStudents)
	api.GET("/students/:id", handlers.GetStudentByID)
	api.PUT("/students/:id", handlers.UpdateStudent)
	api.DELETE("/students/:id", handlers.DeleteStudent)
}