package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/atselvan/go-utils/utils/errors"
	"github.com/stretchr/testify/assert"
)

var (
	testConfigFilePath  = "./config.env"
	mockConfigCreateMsg = "Mock Configuration file %s created"

	testUrl      = "https://test.com"
	testUsername = "test"
	testPassword = "test123"

	testConfigFmt = `MOCK_URL=%s
MOCK_USERNAME=%s
MOCK_PASSWORD=%s`

	testValidConfig      = fmt.Sprintf(testConfigFmt, testUrl, testUsername, testPassword)
	testEmptyConfig      = fmt.Sprintf(testConfigFmt, "", "", "")
	testIncompleteConfig = fmt.Sprintf(testConfigFmt, testUrl, testUsername, "")
)

type (
	MockConfig struct {
		Url      string `mapstructure:"MOCK_URL" required:"true"`
		Username string `mapstructure:"MOCK_USERNAME" required:"true"`
		Password string `mapstructure:"MOCK_PASSWORD" required:"true"`
	}

	MockJsonConfig struct {
		Url      string `json:"MOCK_URL" required:"true"`
		Username string `json:"MOCK_USERNAME" required:"true"`
		Password string `json:"MOCK_PASSWORD" required:"true"`
	}
)

func TestAddConfigPath(t *testing.T) {
	AddConfigPath(testConfigFilePath)
	assert.Equal(t, testConfigFilePath, defaultConfigPath)
}

func TestSetConfigName(t *testing.T) {
	SetConfigName("test")
	assert.Equal(t, "test", defaultConfigName)
}

func TestSetConfigType(t *testing.T) {
	SetConfigType("yaml")
	assert.Equal(t, "yaml", defaultConfigType)
}

func TestLoad(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		resetConfig()
		if err := mockConfig(testConfigFilePath, testValidConfig); err != nil {
			assert.NoError(t, err)
		}
		t.Logf(mockConfigCreateMsg, testConfigFilePath)
		defer removeMockConfig(t, testConfigFilePath)

		config := new(MockConfig)
		cErr := Load(config)
		assert.Nil(t, cErr)
		assert.Equal(t, testUrl, config.Url)
		assert.Equal(t, testUsername, config.Username)
		assert.Equal(t, testPassword, config.Password)
	})

	t.Run("no config file", func(t *testing.T) {
		resetConfig()
		config := &MockConfig{}
		cErr := Load(config)
		assert.NotNil(t, cErr)
		assert.Equal(t, errors.ErrCodeConfigLoad, cErr.Code)
		assert.Contains(t, cErr.Message, "Config File \"config\" Not Found")
	})

	t.Run("elem reflect error", func(t *testing.T) {
		resetConfig()
		cErr := mockConfig(testConfigFilePath, testValidConfig)
		assert.NoError(t, cErr)
		t.Logf(mockConfigCreateMsg, testConfigFilePath)
		defer removeMockConfig(t, testConfigFilePath)

		config := MockConfig{}
		assert.Panics(t, func() {
			_ = Load(config)
		})
	})

	t.Run("empty config", func(t *testing.T) {
		resetConfig()
		if err := mockConfig(testConfigFilePath, testEmptyConfig); err != nil {
			assert.NoError(t, err)
		}
		t.Logf(mockConfigCreateMsg, testConfigFilePath)
		defer removeMockConfig(t, testConfigFilePath)

		t.Run("mapstructure", func(t *testing.T) {
			config := new(MockConfig)
			cErr := Load(config)
			assert.NotNil(t, cErr)
			assert.Equal(t, errors.ErrCodeMissingMandatoryConfiguration, cErr.Code)
			assert.Equal(t, fmt.Sprintf(errors.ErrMsg[errors.ErrCodeMissingMandatoryConfiguration],
				"[MOCK_URL MOCK_USERNAME MOCK_PASSWORD]"), cErr.Message)
		})

		t.Run("json", func(t *testing.T) {
			config := new(MockJsonConfig)
			cErr := Load(config)
			assert.NotNil(t, cErr)
			assert.Equal(t, errors.ErrCodeMissingMandatoryConfiguration, cErr.Code)
			assert.Equal(t, fmt.Sprintf(errors.ErrMsg[errors.ErrCodeMissingMandatoryConfiguration],
				"[MOCK_URL MOCK_USERNAME MOCK_PASSWORD]"), cErr.Message)
		})
	})

	t.Run("incomplete config", func(t *testing.T) {
		resetConfig()
		if err := mockConfig(testConfigFilePath, testIncompleteConfig); err != nil {
			assert.NoError(t, err)
		}
		t.Logf(mockConfigCreateMsg, testConfigFilePath)
		defer removeMockConfig(t, testConfigFilePath)

		config := new(MockConfig)
		cErr := Load(config)
		assert.NotNil(t, cErr)
		assert.Equal(t, errors.ErrCodeMissingMandatoryConfiguration, cErr.Code)
		assert.Equal(t, fmt.Sprintf(errors.ErrMsg[errors.ErrCodeMissingMandatoryConfiguration],
			"[MOCK_PASSWORD]"), cErr.Message)
	})
}

func resetConfig() {
	AddConfigPath(".")
	SetConfigName("config")
	SetConfigType("env")
}

func mockConfig(filePath string, config string) error {
	var (
		file *os.File
		err  error
	)

	// check if file exists
	_, err = os.Stat(filePath)
	if os.IsExist(err) {
		if err := os.Remove(filePath); err != nil {
			return err
		}
	}

	file, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// write default config to the file
	if _, err = file.Write([]byte(config)); err != nil {
		return err
	}

	// Save file changes.
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}

func removeMockConfig(t *testing.T, filePath string) {
	var err = os.Remove(filePath)
	assert.NoError(t, err)
}
