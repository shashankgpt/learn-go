package main

import (
	"codesnooper.com/api/db"
	"codesnooper.com/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
