package usecases_test

import (
	"testing"

	"github.com/RenatoValentim/virtual-bookstore/internal/api/rest/dto"
	"github.com/RenatoValentim/virtual-bookstore/internal/entities"
	"github.com/RenatoValentim/virtual-bookstore/internal/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	dummyID        = uuid.New()
	dummyCreatedAt = "2024-01-01T10:00:00Z"
)

type AuthorDataSpy struct {
	author *entities.Author
}

func (a *AuthorDataSpy) Register(author *entities.Author) error {
	author.ID = dummyID
	author.CreatedAt = dummyCreatedAt
	a.author = author
	return nil
}

func TestRegisterAuthor(t *testing.T) {
	assert := assert.New(t)
	var authorDataFake AuthorDataSpy

	t.Run(`Should register a new author`, func(t *testing.T) {
		input := dto.CreateAuthorInput{
			Name:        "John Doe",
			Email:       "johndoe@example.com",
			Description: "This is a fake description",
		}

		registerAuthor := usecases.NewRegisterAuthor(&authorDataFake)
		err := registerAuthor.Execute(&input)

		assert.Nil(err)
		assert.Equal(authorDataFake.author.ID, dummyID)
		assert.Equal(authorDataFake.author.Name, input.Name)
		assert.Equal(authorDataFake.author.Email, input.Email)
		assert.Equal(authorDataFake.author.Description, input.Description)
		assert.Equal(authorDataFake.author.CreatedAt, dummyCreatedAt)
	})

	t.Run(`Should not register a new author if don't have a email field`, func(t *testing.T) {
		input := dto.CreateAuthorInput{
			Name:        "John Doe",
			Email:       "",
			Description: "This is a fake description",
		}

		registerAuthor := usecases.NewRegisterAuthor(&authorDataFake)
		err := registerAuthor.Execute(&input)

		assert.NotNil(err)
	})

	t.Run(`Should not register a new author if don't have a valid email`, func(t *testing.T) {
		input := dto.CreateAuthorInput{
			Name:        "John Doe",
			Email:       "ç$€§/az@example.com",
			Description: "This is a fake description",
		}

		registerAuthor := usecases.NewRegisterAuthor(&authorDataFake)
		err := registerAuthor.Execute(&input)

		assert.NotNil(err)
	})

	t.Run(`Should not register a new author if don't have a name field`, func(t *testing.T) {
		input := dto.CreateAuthorInput{
			Name:        "",
			Email:       "johndoe@example.com",
			Description: "This is a fake description",
		}

		registerAuthor := usecases.NewRegisterAuthor(&authorDataFake)
		err := registerAuthor.Execute(&input)

		assert.NotNil(err)
	})

	t.Run(`Should not register a new author if don't have a description field`, func(t *testing.T) {
		input := dto.CreateAuthorInput{
			Name:        "John Doe",
			Email:       "johndoe@example.com",
			Description: "",
		}

		registerAuthor := usecases.NewRegisterAuthor(&authorDataFake)
		err := registerAuthor.Execute(&input)

		assert.NotNil(err)
	})

	t.Run(`Should not register a new author if the description field is greater than 400`, func(t *testing.T) {
		desc := `Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore
		culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat
		excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate
		voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia.
		Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem
		duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris
		ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis
		sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.`
		input := dto.CreateAuthorInput{
			Name:        "John Doe",
			Email:       "johndoe@example.com",
			Description: desc,
		}

		registerAuthor := usecases.NewRegisterAuthor(&authorDataFake)
		err := registerAuthor.Execute(&input)

		assert.NotNil(err)
	})
}
