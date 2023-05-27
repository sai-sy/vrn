package main

import (
	"github.com/sai-sy/vrn/controllers"
	"github.com/sai-sy/vrn/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/persons", controllers.findPersons)
	r.GET("/persons/:id", controllers.findPerson)
	r.POST("/persons", controllers.createPerson)
	r.PATCH("/persons/:id", controllers.updatePerson)
	r.DELETE("/persons/:id", controllers.deletePerson)

	// Run the server
	r.Run()
}