package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
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

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error){
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error){
		 _, ok := token.Method.(*jwt.SigningMethodHMAC)
		 if !ok {
			return nil, errors.New("invalid token")			
		 }

		 return []byte(secret_key), nil
	})

	if err != nil{
		return token, err
	}

	return token, nil
}