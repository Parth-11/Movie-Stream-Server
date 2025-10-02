package controllers

import (
	"context"
	"net/http"
	"time"

	database "github.com/Parth-11/Movie-Stream-Server/database"
	model "github.com/Parth-11/Movie-Stream-Server/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies")

func GetMovies() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		var movies []model.Movie

		cursor, err := movieCollection.Find(c, bson.M{})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
			return
		}
		defer cursor.Close(c)

		if err = cursor.All(c, &movies); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies"})
			return
		}

		ctx.JSON(http.StatusOK, movies)
	}
}

func GetMovieByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		movieID := ctx.Param("imdb_id")

		if movieID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Movie ID is required."})
			return
		}

		var movie model.Movie

		err := movieCollection.FindOne(c, bson.M{"imdb_id": movieID}).Decode(&movie)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}

		ctx.JSON(http.StatusOK, movie)
	}
}
