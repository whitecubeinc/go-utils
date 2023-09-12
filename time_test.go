package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDayOfStart(t *testing.T) {
	now := time.Now()
	dayOfStart := time.Now()
	SetDayOfStart(&dayOfStart)

	assert.True(t, now.After(dayOfStart))

	assert.Equal(t, now.Year(), dayOfStart.Year())
	assert.Equal(t, now.Month(), dayOfStart.Month())
	assert.Equal(t, now.Day(), dayOfStart.Day())
	assert.Equal(t, dayOfStart.Hour(), 0)
	assert.Equal(t, dayOfStart.Minute(), 0)
	assert.Equal(t, dayOfStart.Second(), 0)
}

func TestDayOfEnd(t *testing.T) {
	now := time.Now()
	dayOfEnd := time.Now()
	SetDayOfEnd(&dayOfEnd)

	assert.True(t, now.Before(dayOfEnd))

	assert.Equal(t, now.Year(), dayOfEnd.Year())
	assert.Equal(t, now.Month(), dayOfEnd.Month())
	assert.Equal(t, now.Day(), dayOfEnd.Day())
	assert.Equal(t, dayOfEnd.Hour(), 23)
	assert.Equal(t, dayOfEnd.Minute(), 59)
	assert.Equal(t, dayOfEnd.Second(), 59)
}

func TestGetTodayOfStart(t *testing.T) {
	now := time.Now().In(time.UTC)
	todayOfStart := GetTodayOfStart(time.UTC)

	assert.True(t, now.After(todayOfStart))

	assert.Equal(t, now.Year(), todayOfStart.Year())
	assert.Equal(t, now.Month(), todayOfStart.Month())
	assert.Equal(t, now.Day(), todayOfStart.Day())
	assert.Equal(t, todayOfStart.Hour(), 0)
	assert.Equal(t, todayOfStart.Minute(), 0)
	assert.Equal(t, todayOfStart.Second(), 0)
}

func TestGetTodayOfEnd(t *testing.T) {
	now := time.Now().In(time.UTC)
	todayOfEnd := GetTodayOfEnd(time.UTC)

	assert.True(t, now.Before(todayOfEnd))

	assert.Equal(t, now.Year(), todayOfEnd.Year())
	assert.Equal(t, now.Month(), todayOfEnd.Month())
	assert.Equal(t, now.Day(), todayOfEnd.Day())
	assert.Equal(t, todayOfEnd.Hour(), 23)
	assert.Equal(t, todayOfEnd.Minute(), 59)
	assert.Equal(t, todayOfEnd.Second(), 59)
}

func TestKRLocation(t *testing.T) {
	krNow := time.Now().In(KRLocation)
	utcNow := time.Now().In(time.UTC)

	assert.Equal(t, krNow.Unix(), utcNow.Unix())
	assert.NotEqual(t, krNow.String(), utcNow.String())
}
