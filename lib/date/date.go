package date

import "time"

const (
	location = "Europe/Madrid"
	layout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	location,_ := time.LoadLocation(location)

	return time.Now().In(location)
}

func GetNowString() string {
	return GetNow().Format(layout)
}