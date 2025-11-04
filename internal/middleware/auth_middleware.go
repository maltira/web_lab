package middleware

import (
	"net/http"
	"web-lab/internal/dto"
	"web-lab/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("token") // проверка существования токена
		if err != nil {                    // токена нет
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Code: 401, Error: "unauthorized"})
			return
		}

		// токен есть - валидация токена
		userID, userGroup, err := utils.ValidateToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Code: 401, Error: "unauthorized"})
			return
		}
		c.Set("userID", userID)
		c.Set("userGroup", userGroup)
		c.Next()
	}
}
