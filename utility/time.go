package utility

import "time"

func GetDurationUntilMidnight() time.Duration {
	now := time.Now()
	resetTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Add(24 * time.Hour)
	return resetTime.Sub(now)
}
