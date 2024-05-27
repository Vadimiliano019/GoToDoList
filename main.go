package main

import (
	"golang_project/config"
	"golang_project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDatabase()

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
