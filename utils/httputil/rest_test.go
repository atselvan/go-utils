package httputil

import (
	"net/http"
	"testing"

	"github.com/atselvan/go-utils/utils/config"
	"github.com/stretchr/testify/assert"
)

func TestGetServerURL(t *testing.T) {
	sc := &config.ServerConfig{
		Protocol: "https",
		Host:     "company.com",
		Port:     "8080",
	}

	t.Run("https", func(t *testing.T) {
		assert.Equal(t, GetServerURL(sc, "/test").String(), "https://company.com/test")
	})

	sc.Protocol = "http"
	t.Run("http", func(t *testing.T) {
		assert.Equal(t, GetServerURL(sc, "/test").String(), "http://company.com:8080/test")
	})
}

func TestNewStringToJsonResponder(t *testing.T) {
	responder := NewStringToJsonResponder(http.StatusOK, "")
	assert.NotNil(t, responder)
}
