package routes

import (
	"net/http"
	"strconv"

	"deva.com/backend/v2/models"
	"deva.com/backend/v2/utils"
	"github.com/gin-gonic/gin"
)

func getEventById(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if eventID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Event ID is required"})
		return
	}
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Event ID"})
		return
	}

	event, err1 := models.GetEventById(eventID)
	if err1 != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}
	userId, err := utils.ValidateJWT(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.UserID = userId // Assuming userId is int64 and matches Event.UserID type
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if eventID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Event ID is required"})
		return
	}
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Event ID"})
		return
	}
	_, err = models.GetEventById(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = eventID
	err = event.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event" + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": event})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
