package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	// Event routes
	router.GET("/events", getEvents)
	router.POST("/events", createEvent)
	router.GET("/events/:id", getEventById)
	router.PUT("/events/:id", updateEvent)
	router.DELETE("/events/:id", deleteEvent)
	// // User routes
	// router.GET("/users", GetUsers)
	// router.POST("/users", CreateUser)
	// router.GET("/users/:id", GetUserById)

	// // Authentication routes
	router.POST("/login", login)
	router.POST("/register", register)

	// Additional routes can be added here
}
