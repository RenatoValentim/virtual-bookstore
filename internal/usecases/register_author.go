package usecases

import (
	"errors"

	"github.com/RenatoValentim/virtual-bookstore/internal/api/rest/dto"
	"github.com/RenatoValentim/virtual-bookstore/internal/constants"
	db_contracts "github.com/RenatoValentim/virtual-bookstore/internal/db/contracts"
	"github.com/RenatoValentim/virtual-bookstore/internal/entities"
	"github.com/RenatoValentim/virtual-bookstore/internal/validation"
)

type registerAuthor struct {
	authorData db_contracts.AuthorData
}

func NewRegisterAuthor(authorData db_contracts.AuthorData) *registerAuthor {
	return &registerAuthor{
		authorData: authorData,
	}
}

func (ra *registerAuthor) validate(author *dto.CreateAuthorInput) error {
	err := validation.EmailValitation(author.Email)
	if err != nil {
		return err
	}
	if author.Name == "" {
		return errors.New("A Name field is required.")
	}
	if author.Description == "" {
		return errors.New("Description field is required.")
	}
	if len(author.Description) > constants.MAX_LENGTH_CONTENT {
		return errors.New("The description field must have a maximum 400 characters.")
	}
	return nil
}

func (ra *registerAuthor) Execute(input *dto.CreateAuthorInput) error {
	err := ra.validate(input)
	if err != nil {
		return err
	}
	ra.authorData.Register(&entities.Author{
		Name:        input.Name,
		Email:       input.Email,
		Description: input.Description,
	})
	return nil
}
