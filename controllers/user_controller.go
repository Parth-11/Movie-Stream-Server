package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	database "github.com/Parth-11/Movie-Stream-Server/database"
	model "github.com/Parth-11/Movie-Stream-Server/models"
	"github.com/Parth-11/Movie-Stream-Server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection("users")

func HashPassword(password string) (string, error) {
	hash_pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Print("Error encrypting the password")
		return "", err
	}

	return string(hash_pass), nil
}

func RegisterUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
			return
		}

		validate := validator.New()

		if err := validate.Struct(user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
			return
		}

		hashedPassword, err := HashPassword(user.Password)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to hash the password"})
			return
		}

		c, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		count, err := userCollection.CountDocuments(c, bson.M{"email": user.Email})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing user"})
			return
		}

		if count > 0 {
			ctx.JSON(http.StatusConflict, gin.H{"error": "User already exist"})
			return
		}

		user.UserID = bson.NewObjectID().Hex()
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		user.Password = hashedPassword

		result, err := userCollection.InsertOne(c, user)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}

func LoginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userLogin model.UserLogin

		if err := ctx.ShouldBindJSON(&userLogin); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
			return
		}

		c, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var foundUser model.User

		err := userCollection.FindOne(c, bson.M{"email": userLogin.Email}).Decode(&foundUser)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(userLogin.Password))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		token, refreshToken, err := utils.GenerateAllToken(foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.Role, foundUser.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
			return
		}

		err = utils.UpdateAllTokens(foundUser.UserID, token, refreshToken)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the tokens"})
			return
		}

		ctx.JSON(http.StatusOK, model.UserResponse{
			UserID:          foundUser.UserID,
			FirstName:       foundUser.FirstName,
			LastName:        foundUser.LastName,
			Role:            foundUser.Role,
			Email:           foundUser.Email,
			FavouriteGenres: foundUser.FavouriteGenres,
			Token:           foundUser.Token,
			RefreshToken:    foundUser.RefreshToken,
		})
	}
}
