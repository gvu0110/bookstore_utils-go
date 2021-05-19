package date

import "time"

const (
	stringDateLayout = "2006-01-02T15:04:05Z"
	dbDateLayout     = "2006-01-02 15:04:05"
)

// GetNow function returns current UTC time object
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowStringFormat function returns current UTC time in string format
func GetNowStringFormat() string {
	return GetNow().Format(stringDateLayout)
}

// GetNowDBFormat function returns current UTC time in DB format
func GetNowDBFormat() string {
	return GetNow().Format(dbDateLayout)
}
