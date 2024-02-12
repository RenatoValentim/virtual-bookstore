package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/RenatoValentim/virtual-bookstore/internal/api/rest"
)

func MakeHTTPClientSpy(uri, httpMethod string, bodyInput map[string]string) *http.Response {
	routes := rest.LoadRoutes()
	bodyBytes, _ := json.Marshal(bodyInput)
	req := httptest.NewRequest(httpMethod, uri, bytes.NewBuffer(bodyBytes))
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)
	return w.Result()
}
