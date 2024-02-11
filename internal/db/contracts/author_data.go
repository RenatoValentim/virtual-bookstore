package contracts

import "github.com/RenatoValentim/virtual-bookstore/internal/entities"

type AuthorData interface {
	Register(author *entities.Author) error
}
