package user

type Repository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
}
