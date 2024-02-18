package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/RenatoValentim/virtual-bookstore/internal/api/rest/dto"
	"github.com/RenatoValentim/virtual-bookstore/internal/constants"
	"github.com/RenatoValentim/virtual-bookstore/internal/validation"
)

func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	msg := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}
	json.NewEncoder(w).Encode(msg)
}

func AuthorInput(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var author dto.CreateAuthorInput
		err := json.NewDecoder(r.Body).Decode(&author)
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if author.Name == "" || author.Description == "" || author.Email == "" {
			sendErrorResponse(w, http.StatusBadRequest, "To create an author, the fields 'name', 'description' and 'email' are required.")
			return
		}

		if len(author.Description) > constants.MAX_LENGTH_CONTENT {
			sendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("[Description]: %v", errors.New("The description field must have a maximum 400 characters.")))
			return
		}

		err = validation.EmailValitation(author.Email)
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("[Email]: %v", err.Error()))
			return
		}

		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}
