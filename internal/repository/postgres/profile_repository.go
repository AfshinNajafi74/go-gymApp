package postgres

import (
	"github.com/AfshinNajafi74/go-gymApp/internal/domain/profile"
	"gorm.io/gorm"
)

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) profile.Repository {
	return &profileRepository{db: db}
}

// Athlete

func (r *profileRepository) CreateAthleteProfile(p *profile.AthleteProfile) error {
	return r.db.Create(p).Error
}

func (r *profileRepository) GetAthleteProfileByUserID(userID uint) (*profile.AthleteProfile, error) {
	var p profile.AthleteProfile
	err := r.db.Where("user_id = ?", userID).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *profileRepository) UpdateAthleteProfile(p *profile.AthleteProfile) error {
	return r.db.Where("user_id = ?", p.UserID).Updates(p).Error
}

// Coach

func (r *profileRepository) CreateCoachProfile(p *profile.CoachProfile) error {
	return r.db.Create(p).Error
}

func (r *profileRepository) GetCoachProfileByUserID(userID uint) (*profile.CoachProfile, error) {
	var p profile.CoachProfile
	err := r.db.Where("user_id = ?", userID).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *profileRepository) UpdateCoachProfile(p *profile.CoachProfile) error {
	return r.db.Where("user_id = ?", p.UserID).Updates(p).Error
}
