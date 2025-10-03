package main

import (
	"fmt"

	routes "github.com/Parth-11/Movie-Stream-Server/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/echo", func(ctx *gin.Context) { ctx.String(200, "Hello World!") })

	routes.SetupUnprotectedRoute(router)
	routes.SetupProtectedRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to Start server on the port 8080,", err)
	}
}
