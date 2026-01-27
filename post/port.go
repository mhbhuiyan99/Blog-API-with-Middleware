package post

import (
	"blogAPI/domain"
	pstHndlr "blogAPI/rest/handlers/post"
)

type Service interface {
	pstHndlr.Service
}


type PostRepo interface {
	Create(p domain.Post) (*domain.Post, error)
	Get(id int) (*domain.Post, error)
	List(page, limit int64) ([]*domain.Post, error)
	Delete(postID int) error
	Update(p domain.Post) (*domain.Post, error)
	Count() (int64, error)
}