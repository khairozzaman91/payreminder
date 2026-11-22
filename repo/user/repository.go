package user

import (
	"payreminder/domain"
)

type Repository interface {
    Create(user *domain.User) (*domain.User, error)
    GetByID(id string) (*domain.User, error)
    GetByEmail(email string) (*domain.User, error)
    Update(user *domain.User) (*domain.User, error)
    Delete(id string) error
}