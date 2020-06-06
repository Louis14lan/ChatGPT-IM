package util

import (
	"time"
)

const (
	YYYY           = "2006"
	YYYYMM         = "2006-01"
	YYYYMMDD       = "2006-01-02"
	YYYYMMDDHHMM   = "2006-01-02 15:04"
	YYYYMMDDHHMMSS = "2006-01-02 15:04:05"
	HHMMSS         = "15:04:05"
	HHMM           = "15:04"
)

func FormatDatetime(datetime time.Time, pattern string) string {
	return datetime.Format(pattern)
}

func ParseDatetime(datetime string, pattern string) (time.Time, error) {
	return time.Parse(pattern, datetime)
}
