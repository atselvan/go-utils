package dateutil

import "time"

const (
	DateFormat     = "2006-01-02"
	DateTimeFormat = time.RFC3339
)

// Now returns the current data and time in the Europe/Amsterdam location.
func Now() time.Time {
	return time.Now()
}

// GetDateTimeStr returns the formatted date and time.
func GetDateTimeStr(format string) string {
	return Now().Format(format)
}
