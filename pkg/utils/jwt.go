package utils

import (
	"errors"
	"fmt"
	"time"
	config "web-lab/configs"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(userID uuid.UUID, userGroup string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":    userID.String(),
			"user_group": userGroup,
			"exp":        time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString([]byte(config.Cfg.Secret))
	if err != nil {
		return "", errors.New(fmt.Sprintf("Ошибка генерации токена, повторите попытку: %v", err))
	}

	return tokenString, nil
}

func ValidateToken(tokenStr string) (userID uuid.UUID, userGroup string, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Cfg.Secret), nil
	})

	if err != nil || !token.Valid {
		return uuid.Nil, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, ok1 := claims["user_id"].(string)
		userGroup, ok2 := claims["user_group"].(string)
		if ok1 && ok2 {
			userUUID, err := uuid.Parse(userID)
			if err != nil {
				return uuid.Nil, "", err
			}

			return userUUID, userGroup, nil
		}
	}
	return uuid.Nil, "", errors.New("invalid token")
}
