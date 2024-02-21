package logger

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var (
	Sink    *MemorySink
	baseUrl = "https://test.com"
)

// MemorySink implements zap.Sink by writing all messages to a buffer.
type MemorySink struct {
	*bytes.Buffer
}

// Implement Close and Sync as no-ops to satisfy the interface. The Write
// method is provided by the embedded buffer.
func (s *MemorySink) Close() error { return nil }
func (s *MemorySink) Sync() error  { return nil }

func configureMockLogger(logLevel string) {
	// Create a sink instance, and register it with zap for the "memory"
	// protocol.
	Sink = &MemorySink{new(bytes.Buffer)}
	_ = zap.RegisterSink("memory", func(*url.URL) (zap.Sink, error) {
		return Sink, nil
	})

	// Redirect all messages to the MemorySink.
	SetLogger(WithLogLevel(logLevel), WithOutputPaths([]string{"memory://"}))
	logger = logger.WithOptions(zap.WithFatalHook(zapcore.WriteThenPanic))
}

func TestSetLoggerWithConfig(t *testing.T) {
	config := getDefaultZapConfig()
	config.OutputPaths = []string{"stderr"}
	SetLoggerWithConfig(config)
}

func TestSetLogger(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		assert.Equal(t, "info", loggerConfig.Level.String())
		assert.Equal(t, []string{"stdout"}, loggerConfig.OutputPaths)
	})

	t.Run("with options", func(t *testing.T) {
		SetLogger(WithLogLevel(LevelDebug), WithOutputPaths([]string{"stderr"}))
		assert.Equal(t, "debug", loggerConfig.Level.String())
		assert.Equal(t, []string{"stderr"}, loggerConfig.OutputPaths)
	})

	t.Run("override options", func(t *testing.T) {
		SetLogger(WithLogLevel(LevelInfo), WithOutputPaths([]string{"stderr"}))
		assert.Equal(t, "info", loggerConfig.Level.String())
		assert.Equal(t, []string{"stderr"}, loggerConfig.OutputPaths)

		SetLogger(WithLogLevel(LevelDebug))
		assert.Equal(t, "debug", loggerConfig.Level.String())
		assert.Equal(t, []string{"stderr"}, loggerConfig.OutputPaths)

		SetLogger(WithOutputPaths([]string{"stdout"}))
		assert.Equal(t, "debug", loggerConfig.Level.String())
		assert.Equal(t, []string{"stdout"}, loggerConfig.OutputPaths)
	})
}

func TestInfo(t *testing.T) {
	configureMockLogger(LevelInfo)
	Info("info message")

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"info\""))
	assert.True(t, strings.Contains(output, "info message"))
}

func TestInfof(t *testing.T) {
	configureMockLogger(LevelInfo)
	Infof("%s info message", "formatted")

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"info\""))
	assert.True(t, strings.Contains(output, fmt.Sprintf("%s info message", "formatted")))
}

func TestWarn(t *testing.T) {
	configureMockLogger(LevelInfo)
	Warn("warn message")

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"warn\""))
	assert.True(t, strings.Contains(output, "warn message"))
}

func TestWarnf(t *testing.T) {
	configureMockLogger(LevelInfo)
	Warnf("%s warn message", "formatted")

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"warn\""))
	assert.True(t, strings.Contains(output, fmt.Sprintf("%s warn message", "formatted")))
}

func TestError(t *testing.T) {
	configureMockLogger(LevelInfo)
	Error("error message")

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"error\""))
	assert.True(t, strings.Contains(output, "error message"))
}

func TestErrorf(t *testing.T) {
	configureMockLogger(LevelInfo)
	Errorf("%s error message", "formatted")

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"error\""))
	assert.True(t, strings.Contains(output, fmt.Sprintf("%s error message", "formatted")))
}

func TestDebug(t *testing.T) {
	configureMockLogger(LevelDebug)

	Debug("debug message")

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"debug\""))
	assert.True(t, strings.Contains(output, "debug message"))
}

func TestDebugf(t *testing.T) {
	configureMockLogger(LevelDebug)

	Debugf("%s debug message", "formatted")

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"debug\""))
	assert.True(t, strings.Contains(output, fmt.Sprintf("%s debug message", "formatted")))
}

func TestFatal(t *testing.T) {
	configureMockLogger(LevelInfo)

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, "fatal message", r)
		}
	}()

	Fatal("fatal message")
}

func TestFatalf(t *testing.T) {
	configureMockLogger(LevelInfo)

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, "formatted fatal message", r)
		}
	}()

	Fatalf("%s fatal message", "formatted")
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() {
		Panic("panic message")
	})
}

func TestPanicf(t *testing.T) {
	assert.Panics(t, func() {
		Panicf("%s panic message", "formatted")
	})
}

func TestGinZap(t *testing.T) {
	configureMockLogger(LevelInfo)
	r := newRouter()

	apiPath := "/test"
	msg := "some message"
	r.GET(apiPath, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, msg)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, apiPath, nil)
	r.ServeHTTP(w, req)

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"info\""))
	assert.True(t, strings.Contains(output, "\"path\":\"/test\""))
	assert.True(t, strings.Contains(output, "\"status\":200"))
}

func TestGinZapError(t *testing.T) {
	configureMockLogger(LevelInfo)
	r := newRouter()

	apiPath := "/test"
	msg := "some message"
	r.GET(apiPath, func(ctx *gin.Context) {
		ctx.Errors = append(ctx.Errors, &gin.Error{
			Err: errors.New(""),
		})
		ctx.JSON(http.StatusOK, msg)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, apiPath, nil)
	r.ServeHTTP(w, req)

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "\"level\":\"error\""))
	assert.True(t, strings.Contains(output, "\"caller\":\"gin"))
}

func TestRestyDebugLogs(t *testing.T) {
	client := resty.New().SetBaseURL(baseUrl)
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	responder := httpmock.NewStringResponder(http.StatusOK, "someString")
	httpmock.RegisterResponder(http.MethodGet, "/", responder)

	resp, err := client.R().Get("/")
	assert.NoError(t, err)

	configureMockLogger(LevelDebug)

	RestyDebugLogger(resp)

	// Assert sink contents
	output := Sink.String()
	t.Logf("output = %s", output)

	assert.True(t, strings.Contains(output, "Request Url: "+baseUrl))
	assert.True(t, strings.Contains(output, "Request Header: map[Authorization:[]"))
	assert.True(t, strings.Contains(output, "Request Body: <nil>"))
	assert.True(t, strings.Contains(output, "someString"))
}

func newRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(GinZap())
	r.Use(gin.Recovery())
	return r
}
