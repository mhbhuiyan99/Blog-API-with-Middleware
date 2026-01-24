package user

import "blogAPI/domain"

type Service interface {
	Create(user domain.User) (*domain.User, error)
	Find(email string) (*domain.User, error)
}