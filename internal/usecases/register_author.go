package usecases

import (
	"errors"

	db_contracts "github.com/RenatoValentim/virtual-bookstore/internal/db/contracts"
)

type registerAuthor struct {
	authorData db_contracts.AuthorData
}

func NewRegisterAuthor(authorData db_contracts.AuthorData) *registerAuthor {
	return &registerAuthor{
		authorData: authorData,
	}
}

func (ra *registerAuthor) Execute(author *db_contracts.Author) error {
	if author.CreatAt == "" {
		return errors.New("A CreatAt field is required.")
	}
	if author.Email == "" {
		return errors.New("An Email field is required.")
	}
	if author.Name == "" {
		return errors.New("An Name field is required.")
	}
	ra.authorData.Register(author)
	return nil
}
