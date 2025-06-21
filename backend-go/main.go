package main

import (
	"deva.com/backend/v2/db"
	"deva.com/backend/v2/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // Initialize the database connection
	server := gin.Default()
	routes.RegisterRoutes(server) // Register the routes
	server.Run(":8080")
}
