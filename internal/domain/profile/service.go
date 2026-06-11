package profile

import "errors"

type Service interface {
	CreateAthleteProfile(userID uint, height, weight float64, goal Goal, level ExperienceLevel) (*AthleteProfile, error)
	GetAthleteProfile(userID uint) (*AthleteProfile, error)
	UpdateAthleteProfile(p *AthleteProfile) error

	CreateCoachProfile(userID uint, bio, specialization string, yearsOfExp int) (*CoachProfile, error)
	GetCoachProfile(userID uint) (*CoachProfile, error)
	UpdateCoachProfile(p *CoachProfile) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateAthleteProfile(userID uint, height, weight float64, goal Goal, level ExperienceLevel) (*AthleteProfile, error) {
	existing, err := s.repo.GetAthleteProfileByUserID(userID)
	if err == nil && existing != nil {
		return nil, errors.New("athlete profile already exists")
	}

	p := &AthleteProfile{
		UserID:          userID,
		Height:          height,
		Weight:          weight,
		Goal:            goal,
		ExperienceLevel: level,
	}

	if err := s.repo.CreateAthleteProfile(p); err != nil {
		return nil, err
	}

	return p, nil
}

func (s *service) GetAthleteProfile(userID uint) (*AthleteProfile, error) {
	return s.repo.GetAthleteProfileByUserID(userID)
}

func (s *service) UpdateAthleteProfile(p *AthleteProfile) error {
	return s.repo.UpdateAthleteProfile(p)
}

func (s *service) CreateCoachProfile(userID uint, bio, specialization string, yearsOfExp int) (*CoachProfile, error) {
	existing, err := s.repo.GetCoachProfileByUserID(userID)
	if err == nil && existing != nil {
		return nil, errors.New("coach profile already exists")
	}

	p := &CoachProfile{
		UserID:         userID,
		Bio:            bio,
		Specialization: specialization,
		YearsOfExp:     yearsOfExp,
	}

	if err := s.repo.CreateCoachProfile(p); err != nil {
		return nil, err
	}

	return p, nil
}

func (s *service) GetCoachProfile(userID uint) (*CoachProfile, error) {
	return s.repo.GetCoachProfileByUserID(userID)
}

func (s *service) UpdateCoachProfile(p *CoachProfile) error {
	return s.repo.UpdateCoachProfile(p)
}
