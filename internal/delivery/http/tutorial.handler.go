package http

import (
	"net/http"
	"web-lab/internal/dto"
	"web-lab/internal/service"

	"github.com/gin-gonic/gin"
)

type TutorialHandler struct {
	service service.TutorialService
}

func NewTutorialHandler(service service.TutorialService) *TutorialHandler {
	return &TutorialHandler{service: service}
}

func (h *TutorialHandler) GetAll(c *gin.Context) {
	tutorials, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, tutorials)
}
