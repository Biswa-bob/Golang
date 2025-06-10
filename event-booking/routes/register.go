package routes

import (
	"net/http"
	"strconv"

	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func registerForEvents(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID",
		})
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
		})
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not register user",
		})
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Registered!",
	})
}

func cancelRegistrations(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID",
		})
	}

	var event models.Event
	event.ID = eventID

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not cancel registration",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Cancelled!",
	})
}
