package routes

import (
	"github.com/gin-gonic/gin"

	c "golang-studies/go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", c.GetStudents)
	r.GET("/alunos/:name", c.Greeting)

	r.Run()
}
