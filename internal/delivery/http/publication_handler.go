package http

import (
	"net/http"
	"web-lab/internal/dto"
	"web-lab/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PublicationHandler struct {
	sc service.PublicationService
}

func NewPublicationHandler(sc service.PublicationService) *PublicationHandler {
	return &PublicationHandler{sc: sc}
}

func (h *PublicationHandler) CreatePublication(c *gin.Context) {
	var req dto.PublicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Error: "Некорректные данные публикации"})
		return
	}

	err := h.sc.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessfulResponse{Message: "Публикаци успешно создана"})
}

func (h *PublicationHandler) DeletePublication(c *gin.Context) {
	id := c.Param("id")
	publicationID := uuid.MustParse(id)

	err := h.sc.Delete(publicationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessfulResponse{Message: "Публикация удалена"})
}

func (h *PublicationHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	publicationID := uuid.MustParse(id)

	p, err := h.sc.FindByID(publicationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

func (h *PublicationHandler) FindAllPublications(c *gin.Context) {
	publications, err := h.sc.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
	}
	c.JSON(http.StatusOK, publications)
}
