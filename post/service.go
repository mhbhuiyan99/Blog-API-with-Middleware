package post

import "blogAPI/domain"

type service struct {
	pstRepo PostRepo
}

func NewService(pstRepo PostRepo) Service {
	return &service{
		pstRepo: pstRepo,
	}
}

func (svc *service) Create(post domain.Post) (*domain.Post, error) {
	return svc.pstRepo.Create(post)
}

func (svc *service) Get(id int) (*domain.Post, error) {
	return svc.pstRepo.Get(id)
}

func (svc *service) List() ([]*domain.Post, error) {
	return svc.pstRepo.List()
}

func (svc *service) Update(post domain.Post) (*domain.Post, error) {
	return svc.pstRepo.Update(post)
}

func (svc *service) Delete(id int) error {
	return svc.Delete(id)
}
