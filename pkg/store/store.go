package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	DB *mongo.Database
	UC *mongo.Collection
}

func NewStore() *Store {
	db, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://mongo"),
	)

	if err != nil {
		panic(err)
	}

	database := db.Database("redprod-chat")

	return &Store{
		DB: database,
		UC: database.Collection("users"),
	}
}
