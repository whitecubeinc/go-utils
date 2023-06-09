package utils

import "time"

func SetDayOfStart(t *time.Time) {
	*t = time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		0,
		0,
		0,
		0,
		t.Location(),
	)
}

func SetDayOfEnd(t *time.Time) {
	*t = time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		23,
		59,
		59,
		59,
		t.Location(),
	)
}

func GetTodayOfStart(location *time.Location) time.Time {
	now := time.Now().In(location)
	SetDayOfStart(&now)
	return now
}

func GetTodayOfEnd(location *time.Location) time.Time {
	now := time.Now().In(location)
	SetDayOfEnd(&now)
	return now
}
