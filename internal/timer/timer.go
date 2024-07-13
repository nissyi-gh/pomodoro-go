package timer

import "time"

type Timer struct {
	duration time.Duration
}

func NewTimer(duration time.Duration) *Timer {
	return &Timer{
		duration: duration,
	}
}
