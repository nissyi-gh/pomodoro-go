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

func TestTimer_Stop(t *testing.T) {
	duration := 5 * time.Millisecond
	timer := NewTimer(duration)

	timer.Start()
	time.Sleep(2 * time.Millisecond)
	timer.Stop()

	select {
	case <-timer.Done():
		t.Error("Timer should not have finished")
	case <-time.After(4 * time.Millisecond):
		// Test passed.
	}
}

func TestTimer_Done(t *testing.T) {
	duration := 1 * time.Second
	timer := NewTimer(duration)

	timer.Start()

	select {
	case <-timer.Done():
		// Test passed.
	case <-time.After(2 * time.Second):
		t.Error("Timer did not finish within the expected time")
	}
}
