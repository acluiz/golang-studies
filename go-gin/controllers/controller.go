package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	db "golang-studies/go-gin/database"
	m "golang-studies/go-gin/models"
)

func GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, m.Students)
}

func Greeting(c *gin.Context) {
	name := c.Param("name")
	greeting := fmt.Sprintf("Hello, %s!", name)

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"greeting": greeting,
	})
}

func CreateStudent(c *gin.Context) {
	var student m.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	db.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}