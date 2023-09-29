package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Claims struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
	RefreshToken bool `json:"refresh_token"`
}

func CreateOAuth2Tokens(id string, email string, userType string) (accessToken string, refreshToken string, err error) {
	accessClaims := &Claims{
		Id:       id,
		Email:    email,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(), // Продолжительность жизни access токена (30 минут)
		},
	}

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = access.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", err
	}

	refreshClaims := &Claims{
		Id:           id,
		RefreshToken: true, // Пометить как refresh токен
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(), // Продолжительность жизни refresh токена (30 дней)
		},
	}

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refresh.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func VerifyOAuth2Token(token string) (string, string, string, string, error) {
	if token == "" {
		return "", "", "", "", errors.New("token is empty")
	}

	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", "", "", "", errors.New("signature is invalid")
		}
		return "", "", "", "", errors.New("token is invalid")
	}

	if !parsedToken.Valid {
		return "", "", "", "", errors.New("parsed token is invalid")
	}

	if claims.RefreshToken {
		return claims.Id, claims.Email, claims.UserType, "refresh", nil
	}

	return claims.Id, claims.Email, claims.UserType, "access", nil
}
