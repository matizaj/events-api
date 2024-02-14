package main

import (
	"events-api/db"
	"events-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	serverErr := server.Run(":8080")
	if serverErr != nil {
		panic("Server crashed!")
	}

}
