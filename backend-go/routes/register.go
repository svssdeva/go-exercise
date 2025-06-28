package routes

import (
	"net/http"
	"strconv"

	"deva.com/backend/v2/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	// This function will handle the registration for an event
	userId := context.GetInt64("userId") // Get userId from context
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register for event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered for event"})
}

func unregisterFromEvent(context *gin.Context) {

	// This function will handle the unregistration from an event
	userId := context.GetInt64("userId") // Get userId from context
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not unregister from event"})
		return
	}

	err = event.Unregister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not unregister from event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully unregistered from event"})
}
