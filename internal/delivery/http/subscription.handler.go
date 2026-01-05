package http

import (
	"net/http"
	"web-lab/internal/dto"
	"web-lab/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SubscriptionHandler struct {
	sc service.SubscriptionService
}

func NewSubscriptionHandler(sc service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{sc: sc}
}

func (h *SubscriptionHandler) UpdateSubscription(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	tID := c.Param("target_id")
	targetID := uuid.MustParse(tID)

	isSubscribe := c.DefaultQuery("is_subscribe", "false")

	var err error = nil
	if isSubscribe == "true" {
		err = h.sc.Create(userID, targetID)
	} else if isSubscribe == "false" {
		err = h.sc.Delete(userID, targetID)
	} else {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: "Некорректные данные: не указано действие"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessfulResponse{Message: "successfully"})
}

func (h *SubscriptionHandler) GetAllSubscriptions(c *gin.Context) {
	id := c.Param("id")
	userID := uuid.MustParse(id)

	subscribers, err := h.sc.GetUserSubscriptions(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, subscribers)
}

func (h *SubscriptionHandler) GetAllSubscribers(c *gin.Context) {
	id := c.Param("id")
	userID := uuid.MustParse(id)

	subscribers, err := h.sc.GetUserSubscribers(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, subscribers)
}

func (h *SubscriptionHandler) CheckIsSubscribe(c *gin.Context) {
	uID := c.DefaultQuery("user_id", "0")
	tID := c.DefaultQuery("target_id", "0")

	userID, err := uuid.Parse(uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	targetID, err := uuid.Parse(tID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	res := h.sc.CheckIsSubscribed(userID, targetID)
	c.JSON(http.StatusOK, res)
}
