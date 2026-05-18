package util

import "time"

var jakartaLoc *time.Location

func init() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	jakartaLoc = loc
}

// Returns the current time in Jakarta timezone
func GetJakartaTimeNow() time.Time {
	return time.Now().In(jakartaLoc)
}

// Returns the current time in Jakarta timezone
func GetJakartaStartOfDay(t time.Time) time.Time {
	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		0, 0, 0, 0,
		jakartaLoc,
	)
}

// Returns the end of today (23:59:59.999999999)
func GetJakartaEndOfDay(t time.Time) time.Time {
	start := GetJakartaStartOfDay(t)

	return start.Add(24*time.Hour - time.Nanosecond)
}

// Converts a Unix timestamp to Jakarta time
func GetJakartaTimeFromUnix(unix int64) time.Time {
	return time.Unix(unix, 0).In(jakartaLoc)
}
