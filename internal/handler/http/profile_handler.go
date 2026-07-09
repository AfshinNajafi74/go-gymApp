package http

import (
	"encoding/json"
	"net/http"

	"github.com/AfshinNajafi74/go-gymApp/internal/domain/profile"
	"github.com/AfshinNajafi74/go-gymApp/internal/handler/dto"
	"github.com/go-playground/validator/v10"
)

type ProfileHandler struct {
	service profile.Service
}

func NewProfileHandler(s profile.Service) *ProfileHandler {
	return &ProfileHandler{service: s}
}

// CreateAthleteProfile godoc
// @Summary Create athlete profile
// @Description Create a profile for an athlete user
// @Tags profile
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param profile body dto.CreateAthleteProfileRequest true "Athlete Profile Info"
// @Success 201 {object} dto.AthleteProfileResponse
// @Failure 400 {object} dto.ValidationErrorResponse
// @Failure 401 {string} string "Unauthorized"
// @Router /profile/athlete [post]
func (h *ProfileHandler) CreateAthleteProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id")

	var req dto.CreateAthleteProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		writeValidationError(w, err)
		return
	}

	p, err := h.service.CreateAthleteProfile(
		uint(userID.(float64)),
		req.Height,
		req.Weight,
		req.Goal,
		req.ExperienceLevel,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := dto.AthleteProfileResponse{
		UserID:          p.UserID,
		Height:          p.Height,
		Weight:          p.Weight,
		Goal:            p.Goal,
		ExperienceLevel: p.ExperienceLevel,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

// GetAthleteProfile godoc
// @Summary Get athlete profile
// @Tags profile
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.AthleteProfileResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not Found"
// @Router /profile/athlete [get]
func (h *ProfileHandler) GetAthleteProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id")

	p, err := h.service.GetAthleteProfile(uint(userID.(float64)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := dto.AthleteProfileResponse{
		UserID:          p.UserID,
		Height:          p.Height,
		Weight:          p.Weight,
		Goal:            p.Goal,
		ExperienceLevel: p.ExperienceLevel,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

// UpdateAthleteProfile godoc
// @Summary Update athlete profile
// @Tags profile
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param profile body dto.UpdateAthleteProfileRequest true "Updated Athlete Profile"
// @Success 200 {object} dto.AthleteProfileResponse
// @Failure 400 {object} dto.ValidationErrorResponse
// @Failure 401 {string} string "Unauthorized"
// @Router /profile/athlete [put]
func (h *ProfileHandler) UpdateAthleteProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id")

	var req dto.UpdateAthleteProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		writeValidationError(w, err)
		return
	}

	p := &profile.AthleteProfile{
		UserID:          uint(userID.(float64)),
		Height:          req.Height,
		Weight:          req.Weight,
		Goal:            req.Goal,
		ExperienceLevel: req.ExperienceLevel,
	}

	if err := h.service.UpdateAthleteProfile(p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := dto.AthleteProfileResponse{
		UserID:          p.UserID,
		Height:          p.Height,
		Weight:          p.Weight,
		Goal:            p.Goal,
		ExperienceLevel: p.ExperienceLevel,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

// CreateCoachProfile godoc
// @Summary Create coach profile
// @Tags profile
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param profile body dto.CreateCoachProfileRequest true "Coach Profile Info"
// @Success 201 {object} dto.CoachProfileResponse
// @Failure 400 {object} dto.ValidationErrorResponse
// @Failure 401 {string} string "Unauthorized"
// @Router /profile/coach [post]
func (h *ProfileHandler) CreateCoachProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id")

	var req dto.CreateCoachProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		writeValidationError(w, err)
		return
	}

	p, err := h.service.CreateCoachProfile(
		uint(userID.(float64)),
		req.Bio,
		req.Specialization,
		req.YearsOfExp,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := dto.CoachProfileResponse{
		UserID:         p.UserID,
		Bio:            p.Bio,
		Specialization: p.Specialization,
		YearsOfExp:     p.YearsOfExp,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

// GetCoachProfile godoc
// @Summary Get coach profile
// @Tags profile
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.CoachProfileResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not Found"
// @Router /profile/coach [get]
func (h *ProfileHandler) GetCoachProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id")

	p, err := h.service.GetCoachProfile(uint(userID.(float64)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := dto.CoachProfileResponse{
		UserID:         p.UserID,
		Bio:            p.Bio,
		Specialization: p.Specialization,
		YearsOfExp:     p.YearsOfExp,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

// UpdateCoachProfile godoc
// @Summary Update coach profile
// @Tags profile
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param profile body dto.UpdateCoachProfileRequest true "Updated Coach Profile"
// @Success 200 {object} dto.CoachProfileResponse
// @Failure 400 {object} dto.ValidationErrorResponse
// @Failure 401 {string} string "Unauthorized"
// @Router /profile/coach [put]
func (h *ProfileHandler) UpdateCoachProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id")

	var req dto.UpdateCoachProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		writeValidationError(w, err)
		return
	}

	p := &profile.CoachProfile{
		UserID:         uint(userID.(float64)),
		Bio:            req.Bio,
		Specialization: req.Specialization,
		YearsOfExp:     req.YearsOfExp,
	}

	if err := h.service.UpdateCoachProfile(p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := dto.CoachProfileResponse{
		UserID:         p.UserID,
		Bio:            p.Bio,
		Specialization: p.Specialization,
		YearsOfExp:     p.YearsOfExp,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

// writeValidationError is a shared helper to write validation errors consistently
func writeValidationError(w http.ResponseWriter, err error) {
	validationErrors := make(map[string]string)

	for _, fieldErr := range err.(validator.ValidationErrors) {
		switch fieldErr.Tag() {
		case "required":
			validationErrors[fieldErr.Field()] = "this field is required"
		case "gt":
			validationErrors[fieldErr.Field()] = "value must be greater than 0"
		case "gte":
			validationErrors[fieldErr.Field()] = "value must be 0 or greater"
		default:
			validationErrors[fieldErr.Field()] = "invalid value"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(dto.ValidationErrorResponse{
		Message: "validation failed",
		Errors:  validationErrors,
	})
}
