package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	IP     string `json:"ip"`
}

func GenerateToken(userID int, ip string) (string, int64, error) {
	time := time.Now().Add(time.Hour * 32).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time,
		},
		userID,
		ip,
	})

	signedAccessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", 0, err
	}
	return signedAccessToken, time, err
}
func GenerateRefreshToken(userID int, ip string) (string, error) {
	return
}

func HashToken(token string) (string, error) {
	hashToken, err := bcrypt.GenerateFromPassword([]byte(token), 10)
	if err != nil {
		return "", err
	}
	return string(hashToken), nil
}

func CompareHash(token, hashToken string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashToken), []byte(token))
	return err == nil
}

func ExtractPayload(token string) (int, string, error) {
	token = token[len("Bearer "):]
	newToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, "", err
	}
	if claims, ok := newToken.Claims.(tokenClaims); ok && newToken.Valid {
		user_id := claims.UserId
		ip := claims.IP
		return user_id, ip, nil
	}
	return 0, "", fmt.Errorf("Ошибка парсинга jwt")
}
