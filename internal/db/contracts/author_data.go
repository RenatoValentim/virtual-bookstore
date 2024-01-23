package contracts

import "github.com/RenatoValentim/virtual-bookstore/internal/dto"

type AuthorData interface {
	Register(author *dto.Author) error
}
