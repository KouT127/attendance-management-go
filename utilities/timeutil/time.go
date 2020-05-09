package timeutil

import (
	"github.com/KouT127/attendance-management/utilities/timezone"
	"github.com/Songmu/flextime"
	"strconv"
	"time"
)

func GetMonthRange(m int) (time.Time, time.Time, error) {
	t, err := time.Parse("200601", strconv.Itoa(m))
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	location := timezone.JSTLocation()
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, location)
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0).Add(-1 * time.Second)
	return firstOfMonth, lastOfMonth, nil
}

func GetDefaultMonth() (int, error) {
	time := flextime.Now().In(timezone.JSTLocation()).Format("200601")
	month, err := strconv.Atoi(time)
	if err != nil {
		return 0, err
	}
	return month, nil
}
