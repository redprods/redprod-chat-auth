package store

import (
	"context"

	"github.com/redprods/redprod-chat-auth/pkg/pb/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Store) GetUserById(ctx context.Context, id uint32) (*auth.User, error) {
	tx := s.UC.FindOne(ctx, bson.M{
		"id": id,
	})

	user := &auth.User{}

	if err := tx.Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "User not found")
		}
		return nil, err
	}

	return user, nil
}

func (s *Store) FindUser(ctx context.Context, login string) ([]*auth.User, error) {
	cur, err := s.UC.Find(ctx, bson.M{
		"login": login,
	})

	users := []*auth.User{}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return users, nil
		}
	}

	for cur.Next(ctx) {
		user := &auth.User{}
		if err := cur.Decode(user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *Store) CreateUser(ctx context.Context, user *auth.User) (*auth.User, error) {
	tx := s.UC.InsertOne(ctx, bson.M{})
}
