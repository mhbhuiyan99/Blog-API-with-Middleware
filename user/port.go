package user

import (
	"blogAPI/domain"
	userHandler "blogAPI/rest/handlers/user"
)

// get/copy the Service interface from handlers/user/port
type Service interface {
	userHandler.Service // embedding
}

type UserRepo interface {
	Create(u domain.User) (*domain.User, error)
	Get(id int) (*domain.User, error)
	List() ([]domain.User, error)
	Delete(userID int) error
	Update(u domain.User) (*domain.User, error)
	Find(email string) (*domain.User, error)
}