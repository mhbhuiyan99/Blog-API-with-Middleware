package database

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Email	string `json:"email"`
	Password string `json:"password"`
	IsWriter bool `json:"is_writer"`
}

var users []User

func (u User) Store() User {
	if u.ID == "" {
		return u
	}
	users = append(users, u)
	return u
} 

func Find(email, password string) (*User) {
	for _, user := range users {
		if user.Email == email && user.Password == password {
			return &user
		}
	}
	return nil
}