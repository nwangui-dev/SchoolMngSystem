package handlers

import (
	"net/http"

	"github.com/nwangui-dev/SchoolMngSystem/models"
	"github.com/nwangui-dev/SchoolMngSystem/store"

	"github.com/labstack/echo/v4"
)

func CreateStudent(c echo.Context) error {
	student := new(models.Student)
	if err := c.Bind(student); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(student); err != nil {
		return err
	}

	if err := store.GDB.Create(student).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, student)
}

func GetStudents(c echo.Context) error {
	var students []models.Student
	if err := store.GDB.Find(&students).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, students)
}

func GetStudentByID(c echo.Context) error {
	id := c.Param("id")
	var student models.Student

	if err := store.GDB.First(&student, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Student not found"})
	}

	return c.JSON(http.StatusOK, student)
}

func UpdateStudent(c echo.Context) error {
	id := c.Param("id")
	var student models.Student

	if err := store.GDB.First(&student, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Student not found"})
	}

	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	store.GDB.Save(&student)
	return c.JSON(http.StatusOK, student)
}

func DeleteStudent(c echo.Context) error {
	id := c.Param("id")
	if err := store.GDB.Delete(&models.Student{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}