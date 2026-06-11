package profile

import "time"

type Goal string
type ExperienceLevel string

const (
	GoalWeightLoss    Goal = "weight_loss"
	GoalMuscleGain    Goal = "muscle_gain"
	GoalEndurance     Goal = "endurance"
	GoalGeneralHealth Goal = "general_health"

	LevelBeginner     ExperienceLevel = "beginner"
	LevelIntermediate ExperienceLevel = "intermediate"
	LevelAdvanced     ExperienceLevel = "advanced"
)

type AthleteProfile struct {
	ID              uint
	UserID          uint
	Height          float64
	Weight          float64
	Goal            Goal
	ExperienceLevel ExperienceLevel
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
