package main

import (
	"fmt"

	controller "github.com/Parth-11/Movie-Stream-Server/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) { ctx.String(200, "Hello World") })
	router.GET("/movies", controller.GetMovies())
	router.GET("/movie/:imdb_id", controller.GetMovieByID())
	router.POST("/addmovie", controller.AddMovie())
	router.POST("/register", controller.RegisterUser())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to Start server on the port 8080,", err)
	}
}
