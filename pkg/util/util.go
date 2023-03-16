package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/redprods/redprod-chat-auth/pkg/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GenerateJWT(id string) (string, error) {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		Id: id,
	})
	return jwt.SignedString([]byte("secret"))
}

func ParseJWT(token string) (*models.Claims, error) {
	tkn, err := jwt.ParseWithClaims(token,
		&models.Claims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, status.Errorf(codes.Internal, "Unexpected signing method : %v", t.Header["alg"])
			}
			return []byte("secret"), nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := tkn.Claims.(*models.Claims); ok && tkn.Valid {
		return claims, nil
	}

	return nil, status.Errorf(codes.Unauthenticated, "Token is invalid")
}
