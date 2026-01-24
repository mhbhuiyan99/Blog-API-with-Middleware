package post

import "blogAPI/domain"

type Service interface {
	Create (post domain.Post) (*domain.Post, error)
	Delete (id int) (error)
	Update (post domain.Post) (*domain.Post, error)
	Get (id int) (*domain.Post, error)
	List () ([]*domain.Post, error)
}