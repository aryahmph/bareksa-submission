package helper

import "time"

func TimeFormat(time time.Time) string {
	return time.Format("Monday, 02 January 2006")
}
