package dateutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	assert.Equal(t, time.Now().Format(DateFormat), Now().Format(DateFormat))
}

func TestGetDateTime(t *testing.T) {
	assert.Equal(t, time.Now().Format(time.RFC3339), GetDateTimeStr(DateTimeFormat))
}
