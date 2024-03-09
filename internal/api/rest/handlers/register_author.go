package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/RenatoValentim/virtual-bookstore/internal/api/rest/dto"
	"github.com/RenatoValentim/virtual-bookstore/internal/db"
	"github.com/RenatoValentim/virtual-bookstore/internal/usecases"
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
		makeApiResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	makeApiResponse(w, http.StatusOK, "Created")
}
