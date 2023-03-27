package store

import (
	"context"
	"fmt"
	"time"

	"github.com/redprods/redprod-chat-auth/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Store) GetUserById(ctx context.Context, id string) (*models.User, error) {
	user_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	tx := s.UC.FindOne(ctx, bson.M{
		"_id": user_id,
	})

	user := &models.User{}

	if err := tx.Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "User not found")
		}
		return nil, err
	}

	return user, nil
}

func (s *Store) FindUser(ctx context.Context, login string) ([]*models.User, error) {
	cur, err := s.UC.Find(ctx, bson.M{
		"login": bson.M{
			"regex": fmt.Sprintf("%s.*", login),
		},
	})

	users := []*models.User{}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return users, nil
		}
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		user := &models.User{}
		if err := cur.Decode(user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *Store) CreateUser(ctx context.Context, user *models.User) error {
	txf := s.UC.FindOne(ctx, bson.M{
		"login": user.Login,
	})

	err := txf.Err()
	if err == nil {
		return status.Errorf(codes.AlreadyExists, "User with this login already exists")
	}

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}

	tx, err := s.UC.InsertOne(ctx, bson.M{
		"login":      user.Login,
		"password":   user.Password,
		"created_at": time.Now(),
	})

	if err != nil {
		return err
	}

	user.Id = tx.InsertedID.(primitive.ObjectID)

	return nil
}
