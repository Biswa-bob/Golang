package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldnot fetch events try agian later!",
		})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID",
		})
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldnot fetch events try agian later!",
		})
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}
	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldnot create event , Try again later",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Created!",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID",
		})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldnot fetch the event, try agian later!",
		})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorised",
		})
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
	}

	updatedEvent.ID = eventID
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldnot update the event, try agian later!",
		})
		fmt.Println(">>>>>err", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID",
		})
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldnot fetch the event, try agian later!",
		})
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorised to delete event",
		})
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldnot delete the event, try agian later!",
		})
		fmt.Println(">>>>>err", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
