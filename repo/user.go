package repo

import (
	"database/sql"
	"fmt"
	"blogAPI/user"
	"github.com/jmoiron/sqlx"
	"blogAPI/domain"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
func (r userRepo) Create(user domain.User) (*domain.User, error) {
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

func (r userRepo) Find(email string) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT id, username, email, password
		FROM users
		WHERE email=$1
	`
	err := r.db.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	
	return &user, err
}

func (u userRepo) Get(id int) (*domain.User, error) {
	return nil, nil
}

func (u userRepo) List() ([]domain.User, error) {
	return nil, nil
}

func (u userRepo) Delete(userID int) error {
	return nil
}
func (u userRepo) Update(usr domain.User) (*domain.User, error) {
	return nil, nil
}
