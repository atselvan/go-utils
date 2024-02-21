package httputil

import (
	"github.com/atselvan/go-utils/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	TraceIDHeaderKey          = "Trace-Id"
	ConsumerIdHeaderKey       = "Consumer-ID"
	SubjectTokenTypeHeaderKey = "Subject-Token-Type"
	SubjectTokenHeaderKey     = "Subject-Token"

	authenticationSuccessMsg = "Authenticated successfully"
)

func GenerateTraceId(ctx *gin.Context) {
	ctx.Request.Header.Set(TraceIDHeaderKey, uuid.NewString())
	ctx.Next()
}

// BasicAuthRequired is a gin middleware for checking if basic authentication is provided in the request
// The method writes the basic auth to the gin context
// The method returns an errors if basic authentication is not set
func BasicAuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := GetBasicAuthFromHeader(ctx); err != nil {
			BasicAuthError(ctx)
			return
		}
		ctx.Next()
	}
}

// BasicAuth is a gin middleware for validation if basic authentication is provided in the request
// and the auth user and password matches with the stored user accounts.
// The method writes the basic auth to the gin context
// The method returns an errors if basic authentication is not set and
// if the authentication fails to match with a user account.
func BasicAuth(accounts map[string]string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := GetBasicAuthFromHeader(ctx); err != nil {
			BasicAuthError(ctx)
			return
		}

		if val, ok := accounts[ctx.GetString(AuthUserKey)]; ok {
			if val == ctx.GetString(AuthPassKey) {
				logger.Info(authenticationSuccessMsg)
			} else {
				BasicAuthFailed(ctx)
				return
			}
		} else {
			BasicAuthFailed(ctx)
			return
		}
	}
}
