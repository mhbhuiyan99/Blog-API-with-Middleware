package user

import "blogAPI/domain"

type service struct {
	usrRepo UserRepo
}

// returing interface type (Service), so must be implemented
func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	usr, err := svc.usrRepo.Create(user)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil 
	}
	
	return usr, nil
}

func (svc *service) Find(email string) (*domain.User, error) {
	usr, err := svc.usrRepo.Find(email)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil
}


