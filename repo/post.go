package repo

import (
	"blogAPI/domain"
	"blogAPI/post"
	"database/sql"
	"fmt"

	"github.com/gosimple/slug"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)


type PostRepo interface {
	post.PostRepo
}

type postRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) PostRepo {
	return &postRepo{
		db: db,
	}
}

func (r postRepo) Create(p domain.Post) (*domain.Post, error) {

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

func (r postRepo) Get(id int) (*domain.Post, error) {
	var post domain.Post

	query := `
		SELECT
			id,
			user_id, 
			title, 
			content, 
			slug, 
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

func (r postRepo) List() ([]*domain.Post, error) {
	var postsList []*domain.Post

	query := `
		SELECT
			id, 
			user_id,
			title, 
			content, 
			slug,
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

func (r postRepo) Update(p domain.Post) (*domain.Post, error) {

	if p.Title == "" {
		return nil, fmt.Errorf("Title cannot be empty")
	}

	// Update slug if title has changed
	p.Slug = slug.Make(p.Title)

	query := `
		UPDATE posts SET 
			title=$1, 
			content=$2, 
			slug=$3,
			image_url=$4, 
			category=$5, 
			tags=$6 
		WHERE id=$7
	`
	var updated domain.Post
	err := r.db.QueryRow(query,
		p.Title,
		p.Content,
		p.Slug,
		p.ImgURL,
		p.Category,
		pq.Array(p.Tags),
		p.ID,
	).Scan(&updated)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &updated, nil
}
