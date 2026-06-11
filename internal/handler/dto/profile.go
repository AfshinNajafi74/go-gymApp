// internal/handler/dto/profile.go
package dto

import "github.com/AfshinNajafi74/go-gymApp/internal/domain/profile"

// Athlete

type CreateAthleteProfileRequest struct {
	Height          float64                 `json:"height" validate:"required,gt=0"`
	Weight          float64                 `json:"weight" validate:"required,gt=0"`
	Goal            profile.Goal            `json:"goal" validate:"required"`
	ExperienceLevel profile.ExperienceLevel `json:"experience_level" validate:"required"`
}

type UpdateAthleteProfileRequest struct {
	Height          float64                 `json:"height" validate:"required,gt=0"`
	Weight          float64                 `json:"weight" validate:"required,gt=0"`
	Goal            profile.Goal            `json:"goal" validate:"required"`
	ExperienceLevel profile.ExperienceLevel `json:"experience_level" validate:"required"`
}

type AthleteProfileResponse struct {
	UserID          uint                    `json:"user_id"`
	Height          float64                 `json:"height"`
	Weight          float64                 `json:"weight"`
	Goal            profile.Goal            `json:"goal"`
	ExperienceLevel profile.ExperienceLevel `json:"experience_level"`
}

// Coach

type CreateCoachProfileRequest struct {
	Bio            string `json:"bio" validate:"required"`
	Specialization string `json:"specialization" validate:"required"`
	YearsOfExp     int    `json:"years_of_exp" validate:"required,gte=0"`
}

type UpdateCoachProfileRequest struct {
	Bio            string `json:"bio" validate:"required"`
	Specialization string `json:"specialization" validate:"required"`
	YearsOfExp     int    `json:"years_of_exp" validate:"required,gte=0"`
}

type CoachProfileResponse struct {
	UserID         uint   `json:"user_id"`
	Bio            string `json:"bio"`
	Specialization string `json:"specialization"`
	YearsOfExp     int    `json:"years_of_exp"`
}
