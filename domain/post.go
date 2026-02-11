package domain

import (
	"github.com/lib/pq"
	"time"
)

// model or entity >> existence
type Post struct {
	ID        int            `json:"id" db:"id"`
	UserID    int            `json:"user_id" db:"user_id"`
	Title     string         `json:"title" db:"title"`
	Slug      string         `json:"slug" db:"slug"`
	Content   string         `json:"content" db:"content"`
	ImgURL    string         `json:"img_url" db:"image_url"`
	Category  string         `json:"category" db:"category"`
	Tags      pq.StringArray `json:"tags" db:"tags"`
	Published bool           `json:"published" db:"published"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
}

type CountOptions struct {
	UserID    int   
	Category  string
	Tag       string
	Keyword   string
}
