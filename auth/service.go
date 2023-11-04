package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {

}

func NewService() *jwtService {
	return &jwtService{}
}

var secret_key = []byte("TzTIddnRLBncmXiy83YT68e2_4E3mTgvT9Be_H1mXhU")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken,err := token.SignedString(secret_key)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}