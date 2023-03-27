package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Login     string             `json:"login,omitempty"`
	Password  string             `json:"password,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
}
