package repo

import (
	"blogAPI/domain"
	"blogAPI/post"
	"database/sql"
	"fmt"
	"strings"

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

func (r postRepo) List(page, limit int64) ([]*domain.Post, error) {
	
	offset := (page - 1)*limit
	
	var postsList []*domain.Post

	query := `
		SELECT
			*
		FROM posts
		LIMIT $1
		OFFSET $2
		`
	err := r.db.Select(&postsList, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return postsList, nil
}

func (r postRepo) Draft(userId, page, limit int64) ([]*domain.Post, error) {
	
	offset := (page - 1)*limit

	var drafts []*domain.Post

	query := `
		SELECT
			*
		FROM posts
		WHERE user_id = $1 AND published = FALSE
		LIMIT $2
		OFFSET $3
		`
	err := r.db.Select(&drafts, query, userId, limit, offset)
	if err != nil {
		return nil, err
	}
	return drafts, nil
}

func (r postRepo) Published(userId, page, limit int64) ([]*domain.Post, error) {

	offset := (page - 1)*limit

	var publishedPosts []*domain.Post

	query := `
		SELECT
			*
		FROM posts
		WHERE user_id = $1 AND published = TRUE
		LIMIT $2
		OFFSET $3
		`
	err := r.db.Select(&publishedPosts, query, userId, limit, offset)
	if err != nil {
		return nil, err
	}
	return publishedPosts, nil
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
			tags=$6,
			published=$7
		WHERE id=$8
	`
	var updated domain.Post
	err := r.db.QueryRow(query,
		p.Title,
		p.Content,
		p.Slug,
		p.ImgURL,
		p.Category,
		pq.Array(p.Tags),
		p.Published,
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

func (r postRepo) TotalCount(opts domain.CountOptions, isPublished bool) (int64, error) {
	
	query := `SELECT COUNT(*) FROM posts`

	var conditions []string
	var args []interface{}
	argNo := 1

	conditions = append(conditions, fmt.Sprintf("published = $%d", argNo))
	args = append(args, isPublished)
	argNo++

	if opts.UserID != 0 {
		conditions = append(conditions, fmt.Sprintf("user_id = $%d", argNo))
		args = append(args, opts.UserID)
		argNo++
	}

	if opts.Category != "" {
		conditions = append(conditions, fmt.Sprintf("category = $%d", argNo))
		args = append(args, opts.Category)
		argNo++
	}

	if opts.Tag != "" {
		conditions = append(conditions, fmt.Sprintf("$%d = ANY(tags)", argNo))
		args = append(args, opts.Tag)
		argNo++
	}

	if opts.Keyword != "" {
		keyword := "%" + opts.Keyword + "%"
		conditions = append(conditions, fmt.Sprintf("(title ILIKE $%d OR content ILIKE $%d)", argNo, argNo+1))
		args = append(args, keyword, keyword)
		argNo += 2
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	var count int64 
	err := r.db.QueryRow(query, args...).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r postRepo) GetByCategory(categoryID int, page, limit int64) ([]*domain.Post, error) {
	offset := (page - 1)*limit

	var postsList []*domain.Post

	query := `
		SELECT
			*
		FROM posts
		WHERE 
			category = $1 AND published = TRUE
		ORDER BY created_at DESC
		LIMIT $2
		OFFSET $3
	`
	err := r.db.Select(&postsList, query, categoryID, limit, offset)
	if err != nil {
		return nil, err
	}
	return postsList, nil
}

func (r postRepo) GetByTag(tagID int, page, limit int64) ([]*domain.Post, error) {
	offset := (page - 1)*limit

	var postsList []*domain.Post

	query := `
		SELECT
			*
		FROM posts
		WHERE $1 = ANY(tags) AND published = TRUE
		ORDER BY created_at DESC
		LIMIT $2
		OFFSET $3
	`
	err := r.db.Select(&postsList, query, tagID, limit, offset)
	if err != nil {
		return nil, err
	}
	return postsList, nil
}
