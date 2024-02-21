package httputil

import (
	"net/http"

	"github.com/atselvan/go-utils/utils/errors"
	"github.com/atselvan/go-utils/utils/logger"
	"github.com/gin-gonic/gin"
)

// NewRouter returns a new gin router which is configured with some settings for logging,
// auto recovery in case of panics and default handlers for NoRoute and MethodNotAllowed.
func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logger.GinZap())
	r.Use(gin.Recovery())
	r.NoRoute(NoRoute)
	r.HandleMethodNotAllowed = true
	r.NoMethod(MethodNotAllowed)
	return r
}

// Health controller handles requests on the /health endpoint.
func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, RestMsg{Message: http.StatusText(http.StatusOK)})
}

// NoRoute no route controller handles request on endpoints that are not configured
func NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, errors.Error{
		Code:    errors.ErrCodePathNotFound,
		Status:  http.StatusNotFound,
		Message: errors.ErrMsg[errors.ErrCodePathNotFound],
		TraceId: ctx.GetHeader(TraceIDHeaderKey),
	})
}

// MethodNotAllowed method not allowed controller handles request on known endpoints but on methods that are not configured
func MethodNotAllowed(ctx *gin.Context) {
	ctx.JSON(http.StatusMethodNotAllowed, errors.Error{
		Code:    errors.ErrCodeMethodNotAllowed,
		Status:  http.StatusMethodNotAllowed,
		Message: errors.ErrMsg[errors.ErrCodeMethodNotAllowed],
		TraceId: ctx.GetHeader(TraceIDHeaderKey),
	})
}
