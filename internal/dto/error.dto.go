package dto

import "web-lab/internal/entity"

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type SuccessfulResponse struct {
	Message string `json:"message"`
}

type SuccessfulAuthResponse struct {
	Message   string       `json:"message"`
	Token     string       `json:"token"`
	User      entity.User  `json:"user"`
	UserGroup entity.Group `json:"user_group"`
}
