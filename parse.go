package gohelper

import (
	"fmt"
	"time"
)

func ParseDate(dateStr string) (*time.Time, error) {
	formats := []string{
		"2006-01-02",           // YYYY-MM-DD
		"02-01-2006",           // DD-MM-YYYY
		"2/1/2006",             // M/D/YYYY
		"02/01/2006",           // DD/MM/YYYY
		"2006/01/02",           // YYYY/MM/DD
		time.RFC3339,           // Support ISO8601/RFC3339
		"2006-01-02T15:04:05Z", // Support ISO8601 without timezone
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return &t, nil
		}
	}

	return nil, fmt.Errorf("invalid date format: %s", dateStr)
}
