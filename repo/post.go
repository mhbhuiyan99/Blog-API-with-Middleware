package repo

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

type PostRepo interface {
	Create(p Post) (*Post, error)
	Get(id int) (*Post, error)
	List() ([]Post, error)
	Delete(postID int) error
	Update(p Post) (*Post, error)
}

type postRepo struct {
	postList []*Post
}

// constructor or constructor-like function
func NewPostRepo() PostRepo {
	return postRepo{

	}
}

func (p postRepo) Create(po Post) (*Post, error) {
	return nil, nil
}
func (p postRepo) Get(id int) (*Post, error) {
	return nil, nil
}
func (p postRepo) List() ([]Post, error) {
	return nil, nil
}
func (p postRepo) Delete(postID int) error {
	return nil
}
func (p postRepo) Update(po Post) (*Post, error) {
	return nil, nil
}
