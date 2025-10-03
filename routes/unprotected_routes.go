package routes

import (
	controller "github.com/Parth-11/Movie-Stream-Server/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUnprotectedRoute(router *gin.Engine) {

	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
	router.GET("/movies", controller.GetMovies())

}
