package repo

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gosimple/slug"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Post struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	Slug      string    `json:"slug" db:"slug"`
	Content   string    `json:"content" db:"content"`
	ImgURL    string    `json:"img_url" db:"image_url"`
	Category  string    `json:"category" db:"category"`
	Tags      []string  `json:"tags" db:"tags"`
	Published bool      `json:"published" db:"published"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type PostRepo interface {
	Create(p Post) (*Post, error)
	Get(id int) (*Post, error)
	List() ([]*Post, error)
	Delete(postID int) error
	Update(p Post) (*Post, error)
}

type postRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) PostRepo {
	return &postRepo{
		db: db,
	}
}

func (r postRepo) Create(p Post) (*Post, error) {

	// Generate slug from title
	p.Slug = slug.Make(p.Title)

	// Generate image URL if not provided
	if p.ImgURL == "" {
		p.ImgURL = fmt.Sprintf("https://unsplash.com/photos/a-computer-keyboard-sitting-on-top-of-a-wooden-desk-Wyc7vHXfCDQ%s", p.Category)
	}

	query := `
		INSERT INTO posts (
		user_id, 
		title, 
		content, 
		slug, 
		image_url, 
		category, 
		tags, 
		published
		)
		VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8
		)
		RETURNING id, created_at, updated_at
	`
	row := r.db.QueryRow(query,
		p.UserID,
		p.Title,
		p.Content,
		p.Slug,
		p.ImgURL,
		p.Category,
		pq.Array(p.Tags),
		p.Published)

	err := row.Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r postRepo) Get(id int) (*Post, error) {
	var post Post

	query := `
		SELECT
			id, 
			user_id, 
			title, 
			content, 
			image_url, 
			category, 
			tags
		FROM posts
		WHERE id=$1
		`
	err := r.db.Get(&post, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &post, nil
}

func (r postRepo) List() ([]*Post, error) {
	var postsList []*Post

	query := `
		SELECT
			id, 
			user_id,
			title, 
			content, 
			image_url,
			category, 
			tags
		FROM posts
		ORDER BY created_at DESC
		`
	err := r.db.Select(&postsList, query)
	if err != nil {
		return nil, err
	}
	return postsList, nil
}

func (r postRepo) Delete(id int) error {
	query := `
		DELETE FROM posts
		WHERE id=$1
		`
	_, err := r.db.Exec(query, id)
	return err
}

func (r postRepo) Update(p Post) (*Post, error) {
	query := `
		UPDATE posts SET 
			title=$1, 
			content=$2, 
			image_url=$3, 
			category=$4, 
			tags=$5 
		WHERE id=$6
	`
	_, err := r.db.Exec(query,
		p.Title,
		p.Content,
		p.ImgURL,
		p.Category,
		pq.Array(p.Tags),
		p.ID,
	)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
