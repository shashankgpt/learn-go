package routes

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"codesnooper.com/api/models"
	"github.com/gin-gonic/gin"
)

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	fmt.Print(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid event id").Error()})
	}
	event, err := models.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.New("event not found").Error()})
	}
	c.JSON(http.StatusOK, event)
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid request body").Error()})
	}
	event.UserId = c.GetInt64("userId")
	fmt.Println(c.GetInt64("userId"))
	event.Save()
	c.JSON(http.StatusOK, event)
}

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid request body").Error()})
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("unable to save user").Error()})
	}
	c.JSON(http.StatusCreated, user)
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid request body").Error()})
		return
	}
	token, err := models.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("invalid credentials").Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
