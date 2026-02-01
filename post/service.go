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

func (svc *service) List(page, limit int64) ([]*domain.Post, error) {
	return svc.pstRepo.List(page, limit)
}

func (svc *service) Draft(userId, page, limit int64) ([]*domain.Post, error) {
	return svc.pstRepo.Draft(userId, page, limit)
}

func (svc *service) Update(post domain.Post) (*domain.Post, error) {
	return svc.pstRepo.Update(post)
}

func (svc *service) Delete(id int) error {
	return svc.Delete(id)
}

func (svc *service) Count() (int64, error) {
	return svc.pstRepo.Count()
}

func (svc *service) CountDrafts(userId int) (int64, error) {
	return svc.pstRepo.CountDrafts(userId)
}
