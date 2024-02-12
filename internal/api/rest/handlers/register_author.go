package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RenatoValentim/virtual-bookstore/internal/api/rest/dto"
)

func RegisterAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var author dto.CreateAuthorInput
	json.NewDecoder(r.Body).Decode(&author)
}
