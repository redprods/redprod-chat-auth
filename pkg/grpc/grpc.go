package grpc

import (
	"context"

	"github.com/redprods/redprod-chat-auth/pkg/models"
	"github.com/redprods/redprod-chat-auth/pkg/pb/auth"
	"github.com/redprods/redprod-chat-auth/pkg/store"
	"github.com/redprods/redprod-chat-auth/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer
	Store *store.Store
}

func (s AuthService) GetUser(ctx context.Context, req *auth.GetUserRequest) (*auth.User, error) {
	user, err := s.Store.GetUserById(ctx, string(req.Id))
	if err != nil {
		return nil, err
	}

	return &auth.User{
		Id:        []byte(user.Id.Hex()),
		Login:     user.Login,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Unix(),
	}, nil
}

func (s AuthService) Register(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	if len(req.Login) < 4 && len(req.Password) < 4 {
		return nil, status.Errorf(codes.InvalidArgument, "length of login and password must be >= 4")
	}

	user := &models.User{
		Login:    req.Login,
		Password: req.Password,
	}

	if err := s.Store.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	token, err := util.GenerateJWT(user.Id.Hex())
	if err != nil {
		return nil, err
	}

	return &auth.AuthResponse{
		AccessToken: token,
	}, nil
}

func (s AuthService) Auth(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	user, err := s.Store.FindUserByLP(ctx, req.Login, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := util.GenerateJWT(string(user.Id))
	if err != nil {
		return nil, err
	}
	return &auth.AuthResponse{
		AccessToken: token,
	}, nil
}

func (s AuthService) FindUser(ctx context.Context, req *auth.FindUsersRequest) (*auth.FindUsersResponse, error) {
	users, err := s.Store.FindUser(ctx, req.Login)
	if err != nil {
		return nil, err
	}

	_users := []*auth.User{}

	for _, user := range users {
		_users = append(_users,
			&auth.User{
				Id:        []byte(user.Id.Hex()),
				Login:     user.Login,
				Password:  user.Password,
				CreatedAt: user.CreatedAt.Unix(),
			})
	}

	return &auth.FindUsersResponse{
		Users: _users,
	}, nil
}

func (s AuthService) Me(ctx context.Context, in *auth.MeRequest) (*auth.User, error) {
	user_id, err := util.ParseJWT(in.AccessToken)
	if err != nil {
		return nil, err
	}

	user, err := s.Store.GetUserById(ctx, user_id.Id)
	if err != nil {
		return nil, err
	}

	return &auth.User{
		Id:        []byte(user.Id.Hex()),
		Login:     user.Login,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Unix(),
	}, nil
}
