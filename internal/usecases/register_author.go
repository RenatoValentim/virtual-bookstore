package usecases

import (
	"errors"

	db_contracts "github.com/RenatoValentim/virtual-bookstore/internal/db/contracts"
	"github.com/RenatoValentim/virtual-bookstore/internal/validation"
)

var maxLengthContent = 400

type registerAuthor struct {
	authorData db_contracts.AuthorData
}

func NewRegisterAuthor(authorData db_contracts.AuthorData) *registerAuthor {
	return &registerAuthor{
		authorData: authorData,
	}
}

func (ra *registerAuthor) validate(author *db_contracts.Author) error {
	err := validation.DateValidation(author.CreatedAt)
	if err != nil {
		return err
	}
	err = validation.EmailValitation(author.Email)
	if err != nil {
		return err
	}
	if author.Name == "" {
		return errors.New("A Name field is required.")
	}
	if author.Description == "" {
		return errors.New("Description field is required.")
	}
	if len(author.Description) > maxLengthContent {
		return errors.New("The description field must have a maximum 400 characters.")
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
