package http

import (
	"errors"
	"net/http"
	"time"
	"web-lab/internal/dto"
	"web-lab/internal/service"
	"web-lab/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	userService  service.UserService
	groupService service.GroupService
}

func NewAuthHandler(userService service.UserService, groupService service.GroupService) *AuthHandler {
	return &AuthHandler{userService: userService, groupService: groupService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	// Получаем тело
	var req dto.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Error: err.Error()})
		return
	}

	// Получаем пользоваетля по почте
	user, err := h.userService.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Error: "Указан неверный email или пароль"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}

	// Проверяем соответствие пароля и генерируем токен
	isPasswordCorrect := utils.CheckPasswordHash(req.Password, user.Password)
	token := ""
	if !isPasswordCorrect {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Error: "Указан неверный email или пароль"})
		return
	} else if user.IsBlock {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Error: "Пользователь был заблокирован в этом сервисе"})
		return
	} else {
		token, err = utils.GenerateToken(user.ID, user.Group.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
			return
		}
		err := h.userService.Update(&dto.UpdateUserRequest{ID: user.ID, LastVisitTime: time.Now()}) // время последнего захода
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
			return
		}
		c.SetCookie("token", token, 3600*24, "/", "", false, true)
	}
	c.JSON(http.StatusOK, dto.SuccessfulAuthResponse{
		Message:   "authorized",
		User:      *user,
		UserGroup: user.Group,
		Token:     token,
	})
}

func (h *AuthHandler) Registration(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Error: "Переданы некорректные данные"})
		return
	}

	_, err := h.userService.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user, err := h.userService.Create(&req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
				return
			}

			token, err := utils.GenerateToken(user.ID, user.Group.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
				return
			}
			c.SetCookie("token", token, 3600*24, "/", "", false, true)

			c.JSON(http.StatusCreated, dto.SuccessfulAuthResponse{
				Message:   "authorized",
				User:      *user,
				UserGroup: user.Group,
				Token:     token,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Error: "Пользователь с таким email уже существует"})
	return
}

func (h *AuthHandler) AuthStatus(c *gin.Context) {
	tokenStr, _ := c.Cookie("token")
	if tokenStr == "" {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Code: 401, Error: "unauthorized"})
		return
	}

	userID, _, err := utils.ValidateToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Code: 401, Error: "unauthorized"})
		return
	}
	user, err := h.userService.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessfulAuthResponse{
		Message:   "authorized",
		User:      *user,
		UserGroup: user.Group,
		Token:     tokenStr,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	_, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Code: 401, Error: "unauthorized"})
		return
	}
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, dto.SuccessfulResponse{Message: "logged out successfully"})
}
