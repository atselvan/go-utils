package httputil

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	noAuthResponse      = `{"errors":[{"code":"BASIC_AUTH_MISSING","status":401,"message":"Basic authentication is required","traceId":"967ed3d6-33ce-4091-943d-b3a6f8b591be"}]}`
	invalidAuthResponse = `{"errors":[{"code":"CREDENTIALS_INVALID","status":401,"message":"Invalid credentials","traceId":"967ed3d6-33ce-4091-943d-b3a6f8b591be"}]}`
)

func setMockTraceId(ctx *gin.Context) {
	ctx.Request.Header.Set(TraceIDHeaderKey, "967ed3d6-33ce-4091-943d-b3a6f8b591be")
	ctx.Next()
}

func setupMockRouter(middleware ...gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware...)
	r.GET("/login", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	return r
}

func TestGenerateTraceId(t *testing.T) {
	router := setupMockRouter(GenerateTraceId)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
	assert.NotNil(t, req.Header.Get(TraceIDHeaderKey))
}

func TestBasicAuthRequired_Success(t *testing.T) {
	router := setupMockRouter(setMockTraceId, BasicAuthRequired())

	// correct auth
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.SetBasicAuth("admin", "admin")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}

func TestBasicAuthRequired_NoAuth(t *testing.T) {
	router := setupMockRouter(setMockTraceId, BasicAuthRequired())
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, noAuthResponse, w.Body.String())
}

func TestBasicAuthRequired_InvalidAuth(t *testing.T) {

	// invalid basic auth
	router := setupMockRouter(setMockTraceId, BasicAuthRequired())
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.Header.Add(AuthorizationHeaderKey, "YWRtaW4=")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, noAuthResponse, w.Body.String())

	// invalid base64
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/login", nil)
	req.Header.Add(AuthorizationHeaderKey, "Basic YWRtaW4")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, noAuthResponse, w.Body.String())

	// no : in encoded auth
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/login", nil)
	req.Header.Add(AuthorizationHeaderKey, "Basic YWRtaW4=")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, noAuthResponse, w.Body.String())
}

func TestBasicAuth_Success(t *testing.T) {
	router := setupMockRouter(BasicAuth(getMockAccount()))

	// admin
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.SetBasicAuth("admin", "admin")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())

	// user
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/login", nil)
	req.SetBasicAuth("user", "user")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}

func TestBasicAuth_Success_NoAuth(t *testing.T) {
	router := setupMockRouter(setMockTraceId, BasicAuth(getMockAccount()))

	// admin
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, noAuthResponse, w.Body.String())
}

func TestBasicAuth_Success_InvalidAuth(t *testing.T) {
	router := setupMockRouter(setMockTraceId, BasicAuth(getMockAccount()))

	// admin
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.SetBasicAuth("admin", "")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, invalidAuthResponse, w.Body.String())

	// non-existing user
	// admin
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/login", nil)
	req.SetBasicAuth("admi", "admin")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, invalidAuthResponse, w.Body.String())
}
