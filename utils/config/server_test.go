package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testServerConfig = `# Configuration
SERVER_PROTOCOL=http
SERVER_HOST=localhost
SERVER_PORT=8000
SERVER_LOG_LEVEL=INFO
STATIC_FILES_ROOT=static
HTML_TEMPLATE_FILES_ROOT=templates
`

	testServerConfigOnlyRequired = `# Configuration
SERVER_HOST=localhost
SERVER_PORT=8000
`

	testServerConfigMissingRequired = `# Configuration
SERVER_PROTOCOL=http
SERVER_LOG_LEVEL=INFO
STATIC_FILES_ROOT=static
HTML_TEMPLATE_FILES_ROOT=templates
`

	testServerConfigInvalidProtocol = `# Configuration
SERVER_PROTOCOL=htt
SERVER_HOST=localhost
SERVER_PORT=8000
`

	testServerConfigInvalidLogLevel = `# Configuration
SERVER_HOST=localhost
SERVER_PORT=8000
SERVER_LOG_LEVEL=TST
`
)

func TestLoadServerConfig(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		dErr := mockConfig(testConfigFilePath, testServerConfig)
		assert.NoError(t, dErr)
		defer removeMockConfig(t, testConfigFilePath)

		cnf, err := LoadServerConfig()
		assert.Nil(t, err)
		assert.Equal(t, "http", cnf.Protocol)
		assert.Equal(t, "localhost", cnf.Host)
		assert.Equal(t, "8000", cnf.Port)
		assert.Equal(t, "INFO", cnf.LogLevel)
		assert.Equal(t, "static", cnf.StaticFilesRoot)
		assert.Equal(t, "templates", cnf.HTMLTemplateFilesRoot)
	})

	t.Run("only required", func(t *testing.T) {
		dErr := mockConfig(testConfigFilePath, testServerConfigOnlyRequired)
		assert.NoError(t, dErr)
		defer removeMockConfig(t, testConfigFilePath)

		cnf, err := LoadServerConfig()
		assert.Nil(t, err)
		assert.Equal(t, cnf.Protocol, "https")
		assert.Equal(t, cnf.Host, "localhost")
		assert.Equal(t, cnf.Port, "8000")
		assert.Equal(t, cnf.LogLevel, "")
		assert.Equal(t, cnf.StaticFilesRoot, "/")
		assert.Equal(t, cnf.HTMLTemplateFilesRoot, "/")
	})

	t.Run("missing required", func(t *testing.T) {
		dErr := mockConfig(testConfigFilePath, testServerConfigMissingRequired)
		assert.NoError(t, dErr)
		defer removeMockConfig(t, testConfigFilePath)

		cnf, err := LoadServerConfig()
		assert.Nil(t, cnf)
		assert.Equal(t, "Missing mandatory configuration : [SERVER_HOST SERVER_PORT]", err.Message)
	})

	t.Run("invalid protocol", func(t *testing.T) {
		dErr := mockConfig(testConfigFilePath, testServerConfigInvalidProtocol)
		assert.NoError(t, dErr)
		defer removeMockConfig(t, testConfigFilePath)

		cnf, err := LoadServerConfig()
		assert.Nil(t, err)
		assert.Equal(t, cnf.Protocol, "htt")
		assert.Equal(t, cnf.Host, "localhost")
		assert.Equal(t, cnf.Port, "8000")
	})

	t.Run("invalid log level", func(t *testing.T) {
		dErr := mockConfig(testConfigFilePath, testServerConfigInvalidLogLevel)
		assert.NoError(t, dErr)
		defer removeMockConfig(t, testConfigFilePath)

		cnf, err := LoadServerConfig()
		assert.Nil(t, err)
		assert.Equal(t, cnf.Host, "localhost")
		assert.Equal(t, cnf.Port, "8000")
		assert.Equal(t, cnf.LogLevel, "TST")
	})
}
