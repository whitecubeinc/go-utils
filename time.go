package utils

import "time"

var (
	KRLocation = func() *time.Location {
		location, err := time.LoadLocation("Asia/Seoul")
		if err != nil {
			panic(err)
		}
		return location
	}()
)

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

func SetStartOfThisMonth(t *time.Time) {
	*t = time.Date(
		t.Year(),
		t.Month(),
		1,
		0,
		0,
		0,
		0,
		t.Location(),
	)
}

func GetStartOfThisMonth(location *time.Location) time.Time {
	now := time.Now().In(location)
	SetStartOfThisMonth(&now)
	return now
}

func IsSameDate(t1, t2 time.Time, location *time.Location) bool {
	t1 = t1.In(location)
	t2 = t2.In(location)

	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	if y1 == y2 && m1 == m2 && d1 == d2 {
		return true
	}

	return false
}
