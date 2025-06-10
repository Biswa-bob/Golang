package routes

import (
	"example.com/event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) //GET. POST, PUT, PATCH, DELETE
	server.GET("/event/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/event", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvents)
	authenticated.DELETE("/events/:id/delete", cancelRegistrations)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
