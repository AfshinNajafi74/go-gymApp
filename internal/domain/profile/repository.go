package profile

type Repository interface {
	CreateAthleteProfile(p *AthleteProfile) error
	GetAthleteProfileByUserID(userID uint) (*AthleteProfile, error)
	UpdateAthleteProfile(p *AthleteProfile) error

	CreateCoachProfile(p *CoachProfile) error
	GetCoachProfileByUserID(userID uint) (*CoachProfile, error)
	UpdateCoachProfile(p *CoachProfile) error
}
