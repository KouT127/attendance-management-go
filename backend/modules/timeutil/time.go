package timeutil

import (
	"github.com/KouT127/attendance-management/modules/timezone"
	"time"
)

func GetMonthRange(t time.Time) (time.Time, time.Time) {
	year, month, _ := t.Date()
	location := timezone.JSTLocation()

	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, location)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	return firstOfMonth, lastOfMonth
}
