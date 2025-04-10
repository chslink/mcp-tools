package tools

import (
	"time"
)

// ToolGetCurrentTime IANA timezone name (e.g., 'America/New_York', 'Europe/London')
// @timezone test description
func ToolGetCurrentTime(timezone string) (string, error) {
	var tz *time.Location
	if timezone == "" {
		tz = time.Local
	} else {
		tz2, err := time.LoadLocation(timezone)
		if err != nil {
			return "", err
		}
		tz = tz2
	}
	return time.Now().In(tz).Format(time.RFC3339), nil
}

// ToolTest description=xxx
func ToolTest(srcTime string) (string, error) {
	return time.Now().Format(time.RFC3339), nil
}
