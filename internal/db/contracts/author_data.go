package contracts

type Author struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type AuthorData interface {
	Register(author *Author) error
}
