package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/AfshinNajafi74/go-gymApp/internal/domain/user"
)

type UserHandler struct {
	service user.Service
}

func NewUserHandler(s user.Service) *UserHandler {
	return &UserHandler{service: s}
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

	err := h.service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//resp := RegisterResponse{
	//	ID:    user.ID,
	//	Name:  user.Name,
	//	Email: user.Email,
	//}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
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

	// تولید JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := []byte("supersecretkey") // بهتره از cfg.JWT_SECRET بخونی
	tokenString, err := token.SignedString(secret)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	resp := LoginResponse{Token: tokenString}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
