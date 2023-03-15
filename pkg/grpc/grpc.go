package grpc

import (
	"context"

	"github.com/redprods/redprod-chat-auth/pkg/pb/auth"
	"github.com/redprods/redprod-chat-auth/pkg/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer
	Store *store.Store
}

func (s AuthService) GetUser(ctx context.Context, req *auth.GetUserRequest) (*auth.User, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}

func (s AuthService) Register(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}

func (s AuthService) Login(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {

	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}

func (s AuthService) FindUser(ctx context.Context, req *auth.FindUsersRequest) (*auth.FindUsersResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}
