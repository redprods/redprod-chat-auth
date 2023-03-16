package store

import (
	"context"

	"github.com/redprods/redprod-chat-auth/pkg/pb/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Store) FindUserByLP(ctx context.Context, login string, password string) (*auth.User, error) {
	tx := s.UC.FindOne(ctx, bson.M{
		"login":    login,
		"password": password,
	})

	user := &auth.User{}

	if err := tx.Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.Unauthenticated, "Invalid password")
		}
		return nil, err
	}

	return user, nil
}
