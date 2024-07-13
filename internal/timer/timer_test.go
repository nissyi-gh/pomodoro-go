package timer

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	duration := 3 * time.Second
	timer := NewTimer(duration)

	if timer.Duration() != duration {
		t.Errorf("Expected duration to be %v, got %v", duration, timer.duration)
	}
}
