package dto

type RegisterResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type ValidationErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}
