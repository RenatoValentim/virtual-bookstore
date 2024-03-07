package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/RenatoValentim/virtual-bookstore/internal/api/rest/dto"
	"github.com/RenatoValentim/virtual-bookstore/internal/constants"
	"github.com/RenatoValentim/virtual-bookstore/internal/db"
	"github.com/RenatoValentim/virtual-bookstore/internal/usecases"
	"github.com/RenatoValentim/virtual-bookstore/internal/validation"
)

func makeApiResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	msg := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}
	json.NewEncoder(w).Encode(msg)
}

func RegisterAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var author dto.CreateAuthorInput
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		log.Println(err)
		makeApiResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if author.Name == "" || author.Description == "" || author.Email == "" {
		makeApiResponse(w, http.StatusBadRequest, "To create an author, the fields 'name', 'description' and 'email' are required.")
		return
	}

	if len(author.Description) > constants.MAX_LENGTH_CONTENT {
		makeApiResponse(w, http.StatusBadRequest, fmt.Sprintf("[Description]: %v", errors.New("The description field must have a maximum 400 characters.")))
		return
	}

	err = validation.EmailValitation(author.Email)
	if err != nil {
		log.Println(err)
		makeApiResponse(w, http.StatusBadRequest, fmt.Sprintf("[Email]: %v", err.Error()))
		return
	}

	authorData, err := db.NewAuthorDataPostgres()
	if err != nil {
		log.Println(err)
		makeApiResponse(w, http.StatusInternalServerError, "Internal error.")
		return
	}

	registerAuthor := usecases.NewRegisterAuthor(authorData)
	err = registerAuthor.Execute(&author)
	if err != nil {
		log.Println(err)
		makeApiResponse(w, http.StatusInternalServerError, "Failed when try to create a new author.")
		return
	}

	makeApiResponse(w, http.StatusOK, "Created")
}
