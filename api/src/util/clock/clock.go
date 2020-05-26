package clock

import "time"

var fakeTime time.Time

func SetFakeTime(t time.Time) {
	fakeTime = t
}

func ResetFake() {
	fakeTime = time.Time{}
}

func Now() time.Time {
	if !fakeTime.IsZero() {
		return fakeTime
	}
	return time.Now()
}
