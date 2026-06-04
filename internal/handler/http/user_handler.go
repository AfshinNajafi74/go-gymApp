package http

import (
	"encoding/json"
	"net/http"

	"github.com/AfshinNajafi74/go-gymApp/internal/domain/user"
	"github.com/AfshinNajafi74/go-gymApp/internal/handler/dto"
	"github.com/AfshinNajafi74/go-gymApp/pkg/auth"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type UserHandler struct {
	service   user.Service
	jwtSecret string
}

func NewUserHandler(s user.Service, jwtSecret string) *UserHandler {
	return &UserHandler{
		service:   s,
		jwtSecret: jwtSecret,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "User Info"
// @Success 201 {object} RegisterResponse "Created"
// @Failure 400 {string} string "Bad Request"
// @Router /register [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {

		validationErrors := make(map[string]string)

		for _, fieldErr := range err.(validator.ValidationErrors) {

			switch fieldErr.Tag() {

			case "required":
				validationErrors[fieldErr.Field()] = "this field is required"

			case "email":
				validationErrors[fieldErr.Field()] = "invalid email format"

			case "min":
				if fieldErr.Field() == "Password" {
					validationErrors[fieldErr.Field()] =
						"password must be at least 6 characters"
				}

			default:
				validationErrors[fieldErr.Field()] = "invalid value"
			}
		}

		resp := dto.ValidationErrorResponse{
			Message: "validation failed",
			Errors:  validationErrors,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	u, err := h.service.Register(
		req.Name,
		req.Email,
		req.Password,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, err := auth.GenerateToken(u.ID, h.jwtSecret)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	resp := dto.RegisterResponse{
		Token: tokenString,
		User: dto.UserResponse{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Login godoc
// @Summary Login a user
// @Description Authenticate user and return JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param user body LoginRequest true "User Credentials"
// @Success 200 {object} LoginResponse "JWT token returned"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tokenString, err := auth.GenerateToken(user.ID, h.jwtSecret)

	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	resp := LoginResponse{Token: tokenString}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

// Profile godoc
// @Summary Get user profile
// @Tags users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.UserResponse
// @Failure 401 {string} string
// @Router /profile [get]
func (h *UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id")
	u, err := h.service.GetByID(uint(userID.(float64)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	resp := dto.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}
