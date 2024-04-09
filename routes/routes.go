package routes

import (
	"codesnooper.com/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	authenticated :=router.Group("/")
	authenticated.Use(middleware.Authenticate)
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEvent)
	authenticated.POST("/events", createEvent)
	router.POST("/signup", signup)
	router.POST("/login", login)
}
