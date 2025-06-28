package routes

import (
	"deva.com/backend/v2/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Event routes

	authenticated := router.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/events", getEvents)
	authenticated.GET("/events/:id", getEventById)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", unregisterFromEvent)
	// // User routes
	// router.GET("/users", GetUsers)
	// router.POST("/users", CreateUser)
	// router.GET("/users/:id", GetUserById)

	// // Authentication routes
	router.POST("/login", login)
	router.POST("/register", register)

	// Additional routes can be added here
}
