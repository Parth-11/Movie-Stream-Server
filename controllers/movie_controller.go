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
		}
		defer cursor.Close(c)

		if err = cursor.All(c, &movies); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies"})
		}

		ctx.JSON(http.StatusOK, movies)
	}
}
