package middleware

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	"codesnooper.com/api/utils"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errors.New("invalid token").Error()})
		return
	}

	userId, err := hash.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errors.New("unauthorized").Error()})
		return
	}
	c.Set("userId", userId)
	c.Next()
}