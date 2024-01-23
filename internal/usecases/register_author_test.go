package usecases_test

import (
	"testing"

	db_contracts "github.com/RenatoValentim/virtual-bookstore/internal/db/contracts"
	"github.com/RenatoValentim/virtual-bookstore/internal/usecases"
	"github.com/stretchr/testify/assert"
)

type AuthorDataSpy struct {
	author *db_contracts.Author
}

func (a *AuthorDataSpy) Register(author *db_contracts.Author) error {
	a.author = author
	return nil
}

func TestRegisterAuthor(t *testing.T) {
	assert := assert.New(t)

	t.Run(`Should register a new author`, func(t *testing.T) {
		input := db_contracts.Author{
			Name:        "John Doe",
			Email:       "johndoe@example.com",
			Description: "This is a fake description",
			CreatedAt:   "2024-01-01T10:00:00Z",
		}

		var authorDataFake AuthorDataSpy

		registerAuthor := usecases.NewRegisterAuthor(&authorDataFake)
		err := registerAuthor.Execute(&input)

		assert.Nil(err)
		assert.Equal(authorDataFake.author.Name, input.Name)
		assert.Equal(authorDataFake.author.Email, input.Email)
		assert.Equal(authorDataFake.author.Description, input.Description)
		assert.Equal(authorDataFake.author.CreatedAt, input.CreatedAt)
	})

	t.Run(`Should not register a new author if don't have a CreatAt field`, func(t *testing.T) {
		input := db_contracts.Author{
			Name:        "John Doe",
			Email:       "johndoe@example.com",
			Description: "This is a fake description",
			CreatedAt:   "",
		}
		var authorDataFake AuthorDataSpy

		registerAuthor := usecases.NewRegisterAuthor(&authorDataFake)
		err := registerAuthor.Execute(&input)

		assert.NotNil(err)
	})

	t.Run(`Should not register a new author if don't have a valid date`, func(t *testing.T) {
		input := db_contracts.Author{
			Name:        "John Doe",
			Email:       "johndoe@example.com",
			Description: "This is a fake description",
			CreatedAt:   "2024-25-01T10:00:00Z",
		}
		var authorDataFake AuthorDataSpy

		registerAuthor := usecases.NewRegisterAuthor(&authorDataFake)
		err := registerAuthor.Execute(&input)

		assert.NotNil(err)
	})

	t.Run(`Should not register a new author if don't have a email field`, func(t *testing.T) {
		input := db_contracts.Author{
			Name:        "John Doe",
			Email:       "",
			Description: "This is a fake description",
			CreatedAt:   "2024-01-01T10:00:00Z",
		}
		var authorDataFake AuthorDataSpy

		registerAuthor := usecases.NewRegisterAuthor(&authorDataFake)
		err := registerAuthor.Execute(&input)

		assert.NotNil(err)
	})

	t.Run(`Should not register a new author if don't have a name field`, func(t *testing.T) {
		input := db_contracts.Author{
			Name:        "",
			Email:       "johndoe@example.com",
			Description: "This is a fake description",
			CreatedAt:   "2024-01-01T10:00:00Z",
		}
		var authorDataFake AuthorDataSpy

		registerAuthor := usecases.NewRegisterAuthor(&authorDataFake)
		err := registerAuthor.Execute(&input)

		assert.NotNil(err)
	})

	t.Run(`Should not register a new author if don't have a description field`, func(t *testing.T) {
		input := db_contracts.Author{
			Name:        "John Doe",
			Email:       "johndoe@example.com",
			Description: "",
			CreatedAt:   "2024-01-01T10:00:00Z",
		}
		var authorDataFake AuthorDataSpy

		registerAuthor := usecases.NewRegisterAuthor(&authorDataFake)
		err := registerAuthor.Execute(&input)

		assert.NotNil(err)
	})
}
