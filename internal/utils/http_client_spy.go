package utils

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func MakeHTTPClientSpy(uri, httpMethod string, body io.Reader, handler http.HandlerFunc) ([]byte, error) {
	req := httptest.NewRequest(httpMethod, uri, body)
	w := httptest.NewRecorder()
	handler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	return data, err
}
