package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID              bson.ObjectID
	UserID          string
	FirstName       string
	LastName        string
	Email           string
	Password        string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Token           string
	RefreshToken    string
	FavouriteGenres []Genre
}
