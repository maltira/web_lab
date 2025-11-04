package middleware

import (
	"net/http"
	"web-lab/internal/dto"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userGroup := c.GetString("userGroup")
		if userGroup == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Code: 401, Error: "unauthorized"})
			return
		}

		if userGroup != "Админ" {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.ErrorResponse{Code: 403, Error: "Доступ запрещён, требуются права администратора"})
			return
		}
		c.Next()
	}
}
