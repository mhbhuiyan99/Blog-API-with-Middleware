package database

var postList []Post

type Post struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Category string `json:"category"`
	Tags    []string `json:"tags"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func Store(p Post) {
	postList = append(postList, p)
}

func List() []Post {
	return postList
}