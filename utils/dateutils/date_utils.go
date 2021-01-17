package dateutils

import "time"

const (
	apiDateFormat = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

// GetNow return now in UTC time zone
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString return now in string
func GetNowString() string {
	return GetNow().Format(apiDateFormat)
}

// GetNowDBFormat return now in string
func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
