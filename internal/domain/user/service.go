package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(name, email, password string) (*User, error)
	Login(email, password string) (*User, error)
	GetByID(id uint) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Register(name, email, password string) (*User, error) {

	existingUser, err := s.repo.GetByEmail(email)

	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
func (s *service) Login(email, password string) (*User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *service) GetByID(id uint) (*User, error) {
	return s.repo.GetByID(id)
}
