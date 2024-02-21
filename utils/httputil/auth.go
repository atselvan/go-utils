package httputil

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/atselvan/go-utils/utils/errors"
	"github.com/atselvan/go-utils/utils/logger"
	"github.com/gin-gonic/gin"
)

const (
	AuthUserKey = "username"
	AuthPassKey = "password"
)

// GetBasicAuthFromHeader gets basics authentication from the Authorization header.
func GetBasicAuthFromHeader(ctx *gin.Context) *errors.Error {

	if ctx.Request.Header.Get(AuthorizationHeaderKey) == "" {
		return errors.New(
			errors.ErrCodeBasicAuthMissing,
			http.StatusUnauthorized,
			errors.ErrMsg[errors.ErrCodeBasicAuthMissing],
		)
	}

	auth := strings.SplitN(ctx.Request.Header.Get(AuthorizationHeaderKey), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		return errors.New(
			errors.ErrCodeBasicAuthMissing,
			http.StatusUnauthorized,
			errors.ErrMsg[errors.ErrCodeBasicAuthMissing],
		)
	}

	dAuth, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		errors.New(
			errors.ErrCodeBasicAuthMissing,
			http.StatusUnauthorized,
			errors.ErrMsg[errors.ErrCodeBasicAuthMissing],
		)
	}

	cred := strings.SplitN(string(dAuth), ":", 2)

	if len(cred) != 2 {
		return errors.New(
			errors.ErrCodeBasicAuthMissing,
			http.StatusUnauthorized,
			errors.ErrMsg[errors.ErrCodeBasicAuthMissing],
		)
	}

	ctx.Set(AuthUserKey, cred[0])
	ctx.Set(AuthPassKey, cred[1])

	return nil
}

// BasicAuthError writes an errors to the gin context if basic authentication is not provided
func BasicAuthError(ctx *gin.Context) {
	err := errors.New(
		errors.ErrCodeBasicAuthMissing,
		http.StatusUnauthorized,
		errors.ErrMsg[errors.ErrCodeBasicAuthMissing],
	)
	err.TraceId = ctx.GetHeader(TraceIDHeaderKey)
	ctx.JSON(err.Status, errors.Errors{Errors: []errors.Error{*err}})
	logger.Info(err.Message)
	ctx.Abort()
}

// BasicAuthFailed writes a errors to the gin context if basic authentication fails
func BasicAuthFailed(ctx *gin.Context) {
	err := errors.New(
		errors.ErrCodeInvalidCredentials,
		http.StatusUnauthorized,
		errors.ErrMsg[errors.ErrCodeInvalidCredentials],
	)
	err.TraceId = ctx.GetHeader(TraceIDHeaderKey)
	ctx.JSON(err.Status, errors.Errors{Errors: []errors.Error{*err}})
	logger.Info(err.Message)
	ctx.Abort()
}
