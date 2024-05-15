package services

import (
	"errors"
	"fmt"
	"hillel_auction/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
)

func GenerateJWTAccessToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.AccessJWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).UTC().Unix(),
		},
		UserID: userID,
	})

	return token.SignedString([]byte("test"))
}

func GenerateJWTRefreshToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.RefreshJWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).UTC().Unix(),
		},
		UserID: userID,
	})

	return token.SignedString([]byte("test-1"))
}

func GenerateJWTPairTokens(userID int) (*models.JWTPairTokens, error) {
	accessToken, err := GenerateJWTAccessToken(userID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateJWTRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	jwtPair := &models.JWTPairTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return jwtPair, nil
}

func RefreshAccessToken(refreshToken string) (*models.JWTPairTokens, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &models.RefreshJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("test-1"), nil
	})
	if err != nil {
		log.Error("failed to parse token with claims", err)
		return nil, err
	}

	claims, ok := token.Claims.(*models.RefreshJWTClaims)
	if !ok {
		return nil, errors.New("failed to parse claims")
	}

	return GenerateJWTPairTokens(claims.UserID)
}

func ParseAccessToken(accessToken string) (int, error) {
	accessToken = strings.Replace(accessToken, "Bearer ", "", -1)

	token, err := jwt.ParseWithClaims(accessToken, &models.AccessJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("test"), nil
	})
	if err != nil {
		log.Error("failed to parse token with claims", err)
		return 0, err
	}

	claims, ok := token.Claims.(*models.AccessJWTClaims)
	if !ok {
		return 0, errors.New("failed to parse claims")
	}

	return claims.UserID, nil
}
