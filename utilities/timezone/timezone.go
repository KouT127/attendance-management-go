package timezone

import (
	"time"
)

var (
	loc *time.Location
)

func Set(location string) {
	l, err := time.LoadLocation(location)
	if err != nil {
		panic(l)
	}
	loc = l
}

func JSTLocation() *time.Location {
	return loc
}
