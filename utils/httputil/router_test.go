package httputil

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	healthApiPath   = "/health"
	notFoundApiPath = "/notFound"
)

var (
	successResponse          = `{"message":"OK"}`
	pathNotfoundResponse     = `{"code":"PATH_NOT_FOUND","status":404,"message":"Path Not Found","traceId":""}`
	methodNotAllowedResponse = `{"code":"METHOD_NOT_ALLOWED","status":405,"message":"Method Not Allowed","traceId":""}`
)

func readResponseBody(body *bytes.Buffer, t *testing.T) string {
	byt, err := io.ReadAll(body)
	if err != nil {
		t.Error("unexpected test utils errors")
	}
	return string(byt)
}

func TestHealth(t *testing.T) {
	r := NewRouter()
	r.GET(healthApiPath, Health)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, healthApiPath, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, successResponse, readResponseBody(w.Body, t))
}

func TestNoRoute(t *testing.T) {
	r := NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, notFoundApiPath, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, pathNotfoundResponse, readResponseBody(w.Body, t))
}

func TestMethodNotAllowed(t *testing.T) {
	r := NewRouter()
	r.GET(healthApiPath, Health)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, healthApiPath, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, methodNotAllowedResponse, readResponseBody(w.Body, t))
}
