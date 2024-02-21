package logger

import (
	"fmt"
	"log"

	"github.com/atselvan/go-utils/utils/dateutil"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerConfig = getDefaultZapConfig()
	logger       *zap.Logger
	LevelInfo    = zap.InfoLevel.CapitalString()
	LevelDebug   = zap.DebugLevel.CapitalString()
)

type (
	Config struct {
		Level       *zapcore.Level
		OutputPaths []string
	}

	Option func(config *Config)
)

func init() {
	SetLogger()
}

// SetLogger sets the desired loglevel and re-initializes the logger util.
func SetLogger(opts ...Option) {
	config := new(Config)
	for _, opt := range opts {
		opt(config)
	}
	var err error

	if config.Level != nil {
		loggerConfig.Level = zap.NewAtomicLevelAt(*config.Level)
	}

	if config.OutputPaths != nil {
		loggerConfig.OutputPaths = config.OutputPaths
	}

	if logger, err = loggerConfig.Build(); err != nil {
		log.Fatalln("Unable to initialize logger: ", err)
	}
}

// WithLogLevel is an option that can be used to define a custom log level
// Currently only INFO (default) and DEBUG are supported
func WithLogLevel(logLevel string) Option {
	return func(config *Config) {
		var zapLogLevel zapcore.Level
		if logLevel == LevelDebug {
			zapLogLevel = zap.DebugLevel
		} else {
			zapLogLevel = zap.InfoLevel
		}
		config.Level = &zapLogLevel
	}
}

// WithOutputPaths is an option that can be used to define custom log output paths.
// Default path is stdout.
func WithOutputPaths(outputPaths []string) Option {
	return func(config *Config) {
		config.OutputPaths = outputPaths
	}
}

// SetLoggerWithConfig initializes the logger util with a custom configuration.
func SetLoggerWithConfig(config zap.Config) {
	var err error
	if logger, err = config.Build(); err != nil {
		log.Fatalln("Unable to initialize logger: ", err)
	}
}

// getDefaultZapConfig returns the default logger config with the desired log level.
func getDefaultZapConfig() zap.Config {
	return zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "caller",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths: []string{"stdout"},
	}
}

// Info logs a message at the zap.InfoLevel.
// Additional fields can be added to the logger using tags.
func Info(msg string, tags ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, tags...)
}

// Infof logs a formatted message at the zap.InfoLevel.
func Infof(format string, a ...any) {
	logger.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprintf(format, a...))
}

// Warn logs a message at the zap.WarnLevel.
// Additional fields can be added to the logger using tags.
func Warn(msg string, tags ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Warn(msg, tags...)
}

// Warnf logs a formatted message at the zap.WarnLevel.
func Warnf(format string, a ...any) {
	logger.WithOptions(zap.AddCallerSkip(1)).Warn(fmt.Sprintf(format, a...))
}

// Error logs a message at the zap.ErrorLevel.
// Additional fields can be added to the logger using tags.
func Error(msg string, tags ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, tags...)
}

// Errorf logs a formatted message at the zap.ErrorLevel.
func Errorf(format string, a ...any) {
	logger.WithOptions(zap.AddCallerSkip(1)).Error(fmt.Sprintf(format, a...))
}

// Debug logs a message at the zap.DebugLevel.
// Additional fields can be added to the logger using tags.
func Debug(msg string, tags ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Debug(msg, tags...)
}

// Debugf logs a formatted message at the zap.DebugLevel.
func Debugf(format string, a ...any) {
	logger.WithOptions(zap.AddCallerSkip(1)).Debug(fmt.Sprintf(format, a...))
}

// Fatal logs a message at the zap.FatalLevel.
// Additional fields can be added to the logger using tags.
func Fatal(msg string, tags ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Fatal(msg, tags...)
}

// Fatalf logs a formatted message at the zap.FatalLevel.
func Fatalf(format string, a ...any) {
	logger.WithOptions(zap.AddCallerSkip(1)).Fatal(fmt.Sprintf(format, a...))
}

// Panic logs a message at the zap.PanicLevel.
// Additional fields can be added to the logger using tags.
func Panic(msg string, tags ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Panic(msg, tags...)
}

// Panicf logs a formatted message at the zap.PanicLevel.
func Panicf(format string, a ...any) {
	logger.WithOptions(zap.AddCallerSkip(1)).Panic(fmt.Sprintf(format, a...))
}

// GinZap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
func GinZap() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := dateutil.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		end := dateutil.Now()
		latency := end.Sub(start)

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				logger.WithOptions(zap.AddCallerSkip(1)).Error(e)
			}
		} else {
			logger.Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("latency", latency.String()),
				zap.String("trace-id", c.GetHeader("Trace-Id")),
			)
		}
	}
}

// RestyDebugLogger prints debug logs for a http request based on the resty.Response.
func RestyDebugLogger(resp *resty.Response) {
	resp.Request.Header["Authorization"] = []string{}
	Debug(fmt.Sprintf("Request: %v", resp.Request))
	Debug(fmt.Sprintf("Request Url: %v", resp.Request.URL))
	Debug(fmt.Sprintf("Request Header: %v", resp.Request.Header))
	Debug(fmt.Sprintf("Request Body: %v", resp.Request.Body))
	Debug(fmt.Sprintf("Response Status: %v", resp.Status()))
	Debug(fmt.Sprintf("Response Status Code: %v", resp.StatusCode()))
	Debug(fmt.Sprintf("Response Header: %v", resp.Header()))
	Debug(fmt.Sprintf("Response Body: %v", resp.String()))
}
