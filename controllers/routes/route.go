package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin/controllers/people"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/people", people.GetPeople)
	r.GET("/people/:id", people.GetPerson)
	r.POST("/people/:id", people.CreatePerson)
	r.PUT("/people/:id", people.UpdatePerson)
	r.DELETE("/people/:id", people.DeletePerson)

	return r
}
