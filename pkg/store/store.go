package store

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	DB *mongo.Database
	UC *mongo.Collection
}

func NewStore() *Store {
	db, err := mongo.NewClient()
	if err != nil {
		panic(err)
	}

	database := db.Database("redprod-chat")

	return &Store{
		DB: database,
		UC: database.Collection("users"),
	}
}
