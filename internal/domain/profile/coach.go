package profile

import "time"

type CoachProfile struct {
	ID             uint
	UserID         uint
	Bio            string
	Specialization string
	YearsOfExp     int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
