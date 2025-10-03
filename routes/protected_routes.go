package routes

import (
	controller "github.com/Parth-11/Movie-Stream-Server/controllers"
	middleware "github.com/Parth-11/Movie-Stream-Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controller.GetMovieByID())
	router.POST("/addmovie", controller.AddMovie())
}
