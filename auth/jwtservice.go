package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type Service interface {
	GenerateToken(userID int) (signedToken string, err error)
	ValidateToken(signedToken string) (claims *JwtClaim, err error)
}
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaim struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func NewService() *JwtWrapper {
	return &JwtWrapper{}
}

func (j *JwtWrapper) GenerateToken(userID int) (signedToken string, err error) {
	claims := &JwtClaim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return
	}
	return
}

func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(j.SecretKey), nil
	},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Invalid claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token expired")
		return
	}
	return
}
