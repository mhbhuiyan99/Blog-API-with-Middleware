package post

import "blogAPI/domain"

type Service interface {
	Create (post domain.Post) (*domain.Post, error)
	Delete (id int) (error)
	Update (post domain.Post) (*domain.Post, error)
	Get (id int) (*domain.Post, error)
	List (page, limit int64) ([]*domain.Post, error)
	Draft (userId, page, limit int64) ([]*domain.Post, error)
	Published (userId, page, limit int64) ([]*domain.Post, error)
	Count() (int64, error)
	CountUserPost(userId int, isPublished bool) (int64, error)
}