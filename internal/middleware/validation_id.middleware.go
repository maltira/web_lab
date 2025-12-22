package middleware

import (
	"net/http"
	"web-lab/internal/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ValidateUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if _, err := uuid.Parse(id); err != nil {
			// останавливаем цепочку middleware/handlers
			c.AbortWithStatusJSON(http.StatusNotFound, dto.ErrorResponse{Code: 404, Error: "invalid uuid"})
			return
		}
		// UUID валиден - продолжаем выполнение
		c.Next()
	}
}
