package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "golang-studies/go-gin/database"
	m "golang-studies/go-gin/models"
)

func GetStudents(c *gin.Context) {
	var students []m.Student

	db.DB.Find(&students)

	c.JSON(http.StatusOK, students)
}

func GetStudent(c *gin.Context) {
	var student m.Student

	id := c.Param("id")

	db.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "student not found",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var student m.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	db.DB.Create(&student)

	c.JSON(http.StatusCreated, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	db.DB.Delete(&m.Student{}, id)

	c.JSON(http.StatusNoContent, gin.H{})
}