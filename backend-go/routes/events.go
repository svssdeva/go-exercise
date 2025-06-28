package routes

import (
	"net/http"
	"strconv"

	"deva.com/backend/v2/models"
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

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.UserID = context.GetInt64("userId") // Get userId from context
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

	userId := context.GetInt64("userId")

	extractedEvent, err := models.GetEventById(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if extractedEvent.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this event"})
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
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not authorized to delete this event."})
		return
	}

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
