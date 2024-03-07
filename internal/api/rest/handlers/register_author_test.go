package handlers_test

import (
	"net/http"
	"testing"

	"github.com/RenatoValentim/virtual-bookstore/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestAuthorInputMiddlewareValidations(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should return a HTTP status code of 400 Bad Request for an empty body", func(t *testing.T) {
		input := map[string]string{}

		res := utils.MakeHTTPClientSpy("/author/register", http.MethodPost, input)

		assert.Equal(http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Should return a HTTP status code of 400 Bad Request for an greater than 400 description", func(t *testing.T) {
		desc := `Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore
		culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat
		excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate
		voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia.
		Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem
		duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris
		ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis
		sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.`
		input := map[string]string{
			"name":        "John Doe",
			"email":       "johndoe@example.com",
			"description": desc,
		}

		res := utils.MakeHTTPClientSpy("/author/register", http.MethodPost, input)

		assert.Equal(http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Should return a HTTP status code of 400 Bad Request for an invalid email", func(t *testing.T) {
		input := map[string]string{
			"name":        "John Doe",
			"email":       "ç$€§/az@example.com",
			"description": "This is a fake description",
		}

		res := utils.MakeHTTPClientSpy("/author/register", http.MethodPost, input)

		assert.Equal(http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Should return a HTTP status code of 400 Bad Request for an empty email", func(t *testing.T) {
		input := map[string]string{
			"name":        "John Doe",
			"email":       "",
			"description": "This is a fake description",
		}

		res := utils.MakeHTTPClientSpy("/author/register", http.MethodPost, input)

		assert.Equal(http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Should return a HTTP status code of 400 Bad Request for an empty name", func(t *testing.T) {
		input := map[string]string{
			"name":        "",
			"email":       "johndoe@example.com",
			"description": "This is a fake description",
		}

		res := utils.MakeHTTPClientSpy("/author/register", http.MethodPost, input)

		assert.Equal(http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Should return a HTTP status code of 400 Bad Request for an empty description", func(t *testing.T) {
		input := map[string]string{
			"name":        "John Doe",
			"email":       "johndoe@example.com",
			"description": "",
		}

		res := utils.MakeHTTPClientSpy("/author/register", http.MethodPost, input)

		assert.Equal(http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Should return a HTTP status code of 200 OK for a correct input", func(t *testing.T) {
		input := map[string]string{
			"name":        "John Doe",
			"email":       "johndoe@example.com",
			"description": "this is a fake description",
		}

		res := utils.MakeHTTPClientSpy("/author/register", http.MethodPost, input)

		assert.Equal(http.StatusOK, res.StatusCode)
	})
}
