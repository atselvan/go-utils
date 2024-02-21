package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/atselvan/go-utils/utils/errors"
	"github.com/atselvan/go-utils/utils/logger"
	"github.com/atselvan/go-utils/utils/structutil"
	"github.com/spf13/viper"
)

const (
	configLoadSuccessMsg = "Configuration '%s' loaded successfully"
)

var (
	defaultConfigPath = "."
	defaultConfigName = "config"
	defaultConfigType = "env"
)

// AddConfigPath sets a custom config path.
func AddConfigPath(configPath string) {
	defaultConfigPath = configPath
}

// SetConfigName sets a custom config name.
func SetConfigName(configName string) {
	defaultConfigName = configName
}

// SetConfigType sets a custom config type.
func SetConfigType(configType string) {
	defaultConfigType = configType
}

// Load loads the configuration into the input interface.
// The method returns an *errors.Error if:
//   - there is an error while loading the config.
//   - the validation fails.
func Load(c any) *errors.Error {
	viper.AddConfigPath(defaultConfigPath)
	viper.SetConfigName(defaultConfigName)
	viper.SetConfigType(defaultConfigType)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf(errors.ErrMsg[errors.ErrCodeConfigLoad], reflect.TypeOf(c).Elem())
		return errors.New(
			errors.ErrCodeConfigLoad,
			0,
			err.Error(),
		)
	}

	if err := viper.Unmarshal(c); err != nil {
		logger.Errorf(errors.ErrMsg[errors.ErrCodeConfigLoad], reflect.TypeOf(c).Elem())
		return errors.New(
			errors.ErrCodeConfigLoad,
			0,
			err.Error(),
		)
	}

	if cErr := Validate(c); cErr != nil {
		return cErr
	}

	logger.Info(fmt.Sprintf(configLoadSuccessMsg, reflect.TypeOf(c).Elem()))
	return nil
}

// Validate checks if all the required configuration is set.
// Required fields are identified by the struct tag required.
// The method returns an *errors.Error if the required fields are not set.
func Validate(cnf any) *errors.Error {
	var missingParams []string
	sr := reflect.ValueOf(cnf).Elem()

	for i := 0; i < sr.NumField(); i++ {
		if strings.TrimSpace(sr.Field(i).String()) == "" && sr.Type().Field(i).Tag.
			Get(structutil.StructTagRequired) == "true" {
			tagValue := sr.Type().Field(i).Tag.Get(structutil.StructTagMapstructure)
			if tagValue == "" {
				tagValue = sr.Type().Field(i).Tag.Get(structutil.StructTagJson)
			}
			missingParams = append(missingParams, tagValue)
		}
	}
	if len(missingParams) > 0 {
		return errors.Newf(
			errors.ErrCodeMissingMandatoryConfiguration,
			0,
			errors.ErrMsg[errors.ErrCodeMissingMandatoryConfiguration], missingParams,
		)
	}
	return nil
}
