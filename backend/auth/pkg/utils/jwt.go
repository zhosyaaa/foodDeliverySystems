package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/models"
	"time"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours time.Duration
}

type jwtClaims struct {
	jwt.StandardClaims
	Id    uint64
	Email string
}

func (w *JwtWrapper) GenerateToken(user models.User) (string, error) {
	var err error

	claims := &jwtClaims{
		Id:    user.UserId,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(w.ExpirationHours).Unix(),
			Issuer:    w.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (w *JwtWrapper) VerifyToken(signedToken string) (*jwtClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(w.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok {
		return nil, fmt.Errorf("couldn't parse claims")
	}

	// verify token expiration
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}
