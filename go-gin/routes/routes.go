package routes

import (
	"github.com/gin-gonic/gin"

	c "golang-studies/go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", c.GetStudents)
	r.POST("/alunos", c.CreateStudent)
	r.GET("/alunos/:id", c.GetStudent)

	r.Run()
}
