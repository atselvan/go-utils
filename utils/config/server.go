package config

import (
	"github.com/atselvan/go-utils/utils/errors"
	"github.com/atselvan/go-utils/utils/logger"
)

// ServerConfig represents the default server configuration.
type ServerConfig struct {
	Protocol              string `mapstructure:"SERVER_PROTOCOL"`
	Host                  string `mapstructure:"SERVER_HOST" required:"true"`
	Port                  string `mapstructure:"SERVER_PORT" required:"true"`
	LogLevel              string `mapstructure:"SERVER_LOG_LEVEL"`
	StaticFilesRoot       string `mapstructure:"STATIC_FILES_ROOT"`
	HTMLTemplateFilesRoot string `mapstructure:"HTML_TEMPLATE_FILES_ROOT"`
}

// LoadServerConfig loads configuration from the environment and returns a ServerConfig instance.
// If values are not provided for ServerConfig.Host and ServerConfig.Port a error will be returned and
// if values are not provided for  ServerConfig.Protocol and ServerConfig.StaticFilesRoot a default value is set.
// default values:
//
//	ServerConfig.Protocol =  https
//	ServerConfig.StaticFilesRoot = /
//	ServerConfig.HTMLTemplateFilesRoot = /
func LoadServerConfig() (*ServerConfig, *errors.Error) {
	cnf := new(ServerConfig)
	if err := Load(cnf); err != nil {
		return nil, err
	}
	if cnf.Protocol == "" {
		cnf.Protocol = "https"
	}
	if cnf.LogLevel != "" {
		logger.SetLogger(logger.WithLogLevel(cnf.LogLevel))
	}
	if cnf.StaticFilesRoot == "" {
		cnf.StaticFilesRoot = "/"
	}
	if cnf.HTMLTemplateFilesRoot == "" {
		cnf.HTMLTemplateFilesRoot = "/"
	}
	return cnf, nil
}
