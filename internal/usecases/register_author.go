package usecases

import (
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
	ra.authorData.Register(author)
	return nil
}
