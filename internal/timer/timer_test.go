package timer

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	duration := 3 * time.Second
	timer := NewTimer(duration)

	if timer.duration != duration {
		t.Errorf("Expected duration to be %v, got %v", duration, timer.duration)
	}
}

func TestTimer_Start(t *testing.T) {
	duration := 5 * time.Millisecond
	timer := NewTimer(duration)

	startTime := time.Now()
	timer.Start()

	<-timer.Done()
	elapsedTime := time.Since(startTime)

	if elapsedTime < duration {
		t.Errorf("Expected elapsed time to be at least %v, got %v", duration, elapsedTime)
	}
}
