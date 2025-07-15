package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

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