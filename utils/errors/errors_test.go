package errors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	testErrCode = "TEST_ERROR"
	testErrMsg  = "This is a test error"
	testErrFmt  = "The error is: %s"
)

func TestNewError(t *testing.T) {
	err := New(testErrCode, 0, testErrMsg)
	assert.Equal(t, testErrCode, err.Code)
	assert.Equal(t, 0, err.Status)
	assert.Equal(t, testErrMsg, err.Message)
}

func TestNewErrorf(t *testing.T) {
	err := Newf(testErrCode, 0, testErrFmt, testErrMsg)
	assert.Equal(t, testErrCode, err.Code)
	assert.Equal(t, 0, err.Status)
	assert.Equal(t, fmt.Sprintf(testErrFmt, testErrMsg), err.Message)
}

func TestBadRequestError(t *testing.T) {
	err := BadRequestError(testErrMsg)
	assert.Equal(t, ErrCodeBadRequest, err.Code)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, testErrMsg, err.Message)
}

func TestBadRequestErrorf(t *testing.T) {
	err := BadRequestErrorf(testErrFmt, testErrMsg)
	assert.Equal(t, ErrCodeBadRequest, err.Code)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, fmt.Sprintf(testErrFmt, testErrMsg), err.Message)
}

func TestUnauthorizedError(t *testing.T) {
	err := UnauthorizedError(testErrMsg)
	assert.Equal(t, ErrCodeUnauthorized, err.Code)
	assert.Equal(t, http.StatusUnauthorized, err.Status)
	assert.Equal(t, testErrMsg, err.Message)
}

func TestUnauthorizedErrorf(t *testing.T) {
	err := UnauthorizedErrorf(testErrFmt, testErrMsg)
	assert.Equal(t, ErrCodeUnauthorized, err.Code)
	assert.Equal(t, http.StatusUnauthorized, err.Status)
	assert.Equal(t, fmt.Sprintf(testErrFmt, testErrMsg), err.Message)
}

func TestForbiddenError(t *testing.T) {
	err := ForbiddenError(testErrMsg)
	assert.Equal(t, ErrCodeInsufficientAccess, err.Code)
	assert.Equal(t, http.StatusForbidden, err.Status)
	assert.Equal(t, testErrMsg, err.Message)
}

func TestForbiddenErrorf(t *testing.T) {
	err := ForbiddenErrorf(testErrFmt, testErrMsg)
	assert.Equal(t, ErrCodeInsufficientAccess, err.Code)
	assert.Equal(t, http.StatusForbidden, err.Status)
	assert.Equal(t, fmt.Sprintf(testErrFmt, testErrMsg), err.Message)
}

func TestNotFoundError(t *testing.T) {
	err := NotFoundError(testErrMsg)
	assert.Equal(t, ErrCodeNotFound, err.Code)
	assert.Equal(t, http.StatusNotFound, err.Status)
	assert.Equal(t, testErrMsg, err.Message)
}

func TestNotFoundErrorf(t *testing.T) {
	err := NotFoundErrorf(testErrFmt, testErrMsg)
	assert.Equal(t, ErrCodeNotFound, err.Code)
	assert.Equal(t, http.StatusNotFound, err.Status)
	assert.Equal(t, fmt.Sprintf(testErrFmt, testErrMsg), err.Message)
}

func TestConflictError(t *testing.T) {
	err := ConflictError(testErrMsg)
	assert.Equal(t, ErrCodeConflict, err.Code)
	assert.Equal(t, http.StatusConflict, err.Status)
	assert.Equal(t, testErrMsg, err.Message)
}

func TestConflictErrorf(t *testing.T) {
	err := ConflictErrorf(testErrFmt, testErrMsg)
	assert.Equal(t, ErrCodeConflict, err.Code)
	assert.Equal(t, http.StatusConflict, err.Status)
	assert.Equal(t, fmt.Sprintf(testErrFmt, testErrMsg), err.Message)
}

func TestInternalServerError(t *testing.T) {
	err := InternalServerError(testErrMsg)
	assert.Equal(t, ErrCodeInternalServerError, err.Code)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, testErrMsg, err.Message)
}

func TestInternalServerErrorf(t *testing.T) {
	err := InternalServerErrorf(testErrFmt, testErrMsg)
	assert.Equal(t, ErrCodeInternalServerError, err.Code)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, fmt.Sprintf(testErrFmt, testErrMsg), err.Message)
}
