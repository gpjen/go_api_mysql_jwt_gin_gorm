package helper

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

type JWTservice interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTservice {
	return &jwtService{
		secretKey: os.Getenv("JWT_SECRETKEY"),
		issuer:    "agan",
	}
}

func (j *jwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err.Error())
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
