package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Nickname string             `json:"nickname"`
	Slug     string             `json:"slug"`
	Photo    string             `json:"photo"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Phone    string             `json:"phone"`
}
