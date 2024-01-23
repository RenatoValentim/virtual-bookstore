package usecases

import (
	"errors"
	"time"

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

func (ra *registerAuthor) validate(author *db_contracts.Author) error {
	if author.CreatedAt == "" {
		return errors.New("A CreatAt field is required.")
	}
	_, err := time.Parse(time.RFC3339, author.CreatedAt)
	if err != nil {
		return err
	}
	if author.Email == "" {
		return errors.New("An Email field is required.")
	}
	if author.Name == "" {
		return errors.New("A Name field is required.")
	}
	if author.Description == "" {
		return errors.New("A Description field is required.")
	}
	return nil
}

func (ra *registerAuthor) Execute(author *db_contracts.Author) error {
	err := ra.validate(author)
	if err != nil {
		return err
	}
	ra.authorData.Register(author)
	return nil
}
