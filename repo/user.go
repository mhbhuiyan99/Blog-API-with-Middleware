package repo

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	IsWriter  bool      `json:"is_writer" db:"is_writer"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
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
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
func (r userRepo) Create(user User) (*User, error) {
	query := `
		INSERT INTO users (
			username,
			email,
			password,
			is_writer
		)
		VALUES (
			:username,
			:email,
			:password,
			:is_writer
		)
		RETURNING id, created_at, updated_at
	`
	// Execute named query
	var userID int
	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&userID, &user.CreatedAt, &user.UpdatedAt); err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("failed to retrieve inserted user ID")
	}

	user.ID = userID
	return &user, nil
}

func (r userRepo) Find(email, password string) (*User, error) {
	var user User
	query := `
		SELECT id, username, email, password, is_writer
		FROM users
		WHERE email=$1 AND password=$2
		LIMIT 1
	`
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, err
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
