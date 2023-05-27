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
	r.GET("/persons", controllers.Find)
	r.GET("/persons/:id", controllers.FindBook)
	r.POST("/persons", controllers.CreateBook)
	r.PATCH("/persons/:id", controllers.UpdateBook)
	r.DELETE("/persons/:id", controllers.DeleteBook)

	// Run the server
	r.Run()
}