package repo

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsWriter bool   `json:"is_writer"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Get(id int) (*User, error)
	List() ([]User, error)
	Delete(userID int) error
	Update(u User) (*User, error)
	Find(email, password string) (*User, error)
}


type userRepo struct {
	users []*User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (u userRepo) Find(email, password string) (*User, error) {

	return nil, nil
}

func (u userRepo) Create(usr User) (*User, error) {
	return nil, nil
}

func (u userRepo) Get(id int) (*User, error) {
	return nil, nil
}

func (u userRepo) List() ([]User, error) {
	return nil, nil
}

func (u userRepo) Delete(userID int) error {
	return nil
}
func (u userRepo) Update(usr User) (*User, error) {
	return nil, nil
}
