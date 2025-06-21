package main

import (
	"net/http"

	"deva.com/backend/v2/db"
	"deva.com/backend/v2/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // Initialize the database connection
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	// Simulate fetching events from a database or service
	// events := gin.H{
	// 	"events": []gin.H{
	// 		{"id": 1, "name": "Event 1"},
	// 		{"id": 2, "name": "Event 2"},
	// 		{"id": 3, "name": "Event 3"},
	// 	},
	// }

	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = 1     // Simple ID assignment logic
	event.UserID = 1 // Assuming a static user ID for simplicity
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
