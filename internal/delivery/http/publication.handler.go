package http

import (
	"net/http"
	"web-lab/internal/dto"
	"web-lab/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PublicationHandler struct {
	sc          service.PublicationService
	userService service.UserService
}

func NewPublicationHandler(sc service.PublicationService, userService service.UserService) *PublicationHandler {
	return &PublicationHandler{sc: sc, userService: userService}
}

func (h *PublicationHandler) CreatePublication(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)
	user, err := h.userService.GetByID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: "Не удалось получить информацию о пользователе"})
		return
	}
	if !user.IsBlock && user.Group.CanPublishPosts {
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
		return
	}
	c.JSON(http.StatusForbidden, dto.ErrorResponse{Code: 403, Error: "Вам запрещено публиковать посты"})
}

func (h *PublicationHandler) DeletePublication(c *gin.Context) {
	id := c.Param("id")
	publicationID := uuid.MustParse(id)
	publication, err := h.sc.FindByID(publicationID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: "Не удалось получить информацию о публикации"})
		return
	}

	userID := c.MustGet("userID").(uuid.UUID)
	user, err := h.userService.GetByID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: "Не удалось получить информацию о пользователе"})
		return
	}

	if !user.IsBlock && user.Group.CanPublishPosts && (user.ID == publication.UserID || user.Group.Name == "Админ") {
		err := h.sc.Delete(publicationID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.SuccessfulResponse{Message: "Публикация удалена"})
		return
	}
	c.JSON(http.StatusForbidden, dto.ErrorResponse{Code: 403, Error: "Вам запрещено удалять этот пост"})
}

func (h *PublicationHandler) UpdatePublication(c *gin.Context) {
	var req dto.PublicationUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Error: "Некорректные данные публикации"})
		return
	}
	publication, err := h.sc.FindByID(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: "Не удалось получить информацию о публикации"})
		return
	}

	userID := c.MustGet("userID").(uuid.UUID)
	user, err := h.userService.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: "Не удалось получить информацию о пользователе"})
		return
	}

	if !user.IsBlock && user.Group.CanPublishPosts && (user.ID == publication.UserID || user.Group.Name == "Админ") {
		err := h.sc.Update(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: "Не удалось обновить публикацию"})
			return
		}
		c.JSON(http.StatusOK, dto.SuccessfulResponse{Message: "Публикация изменена"})
		return
	}
	c.JSON(http.StatusForbidden, dto.ErrorResponse{Code: 403, Error: "Вам запрещено изменять этот пост"})
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

func (h *PublicationHandler) FindByUserID(c *gin.Context) {
	id := c.Param("id")
	isDraft := c.DefaultQuery("is_draft", "false")
	userID := uuid.MustParse(id)

	publications, err := h.sc.FindByUserID(userID, isDraft == "true")
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
	}
	c.JSON(http.StatusOK, publications)
}

func (h *PublicationHandler) FindAllPublications(c *gin.Context) {
	isDraft := c.DefaultQuery("is_draft", "false")

	publications, err := h.sc.FindAll(isDraft == "true")
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
	}
	c.JSON(http.StatusOK, publications)
}

func (h *PublicationHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.sc.GetAllCategories()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
	}
	c.JSON(http.StatusOK, categories)
}
