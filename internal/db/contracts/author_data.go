package contracts

type Author struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	CreatAt     string `json:"creat_at"`
}

type AuthorData interface {
	Register(author *Author) error
}
