package dto

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type SuccessfulResponse struct {
	Message string `json:"message"`
}

type SuccessfulAuthResponse struct {
	Message   string `json:"message"`
	Token     string `json:"token"`
	UserID    string `json:"user_id"`
	UserGroup string `json:"user_group"`
}
