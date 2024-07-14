package pomodoro

import (
	"testing"
	"time"
)

func TestNewPomodoro(t *testing.T) {
	workDuration := 25 * time.Minute
	shortBreak := 5 * time.Minute
	longBreak := 15 * time.Minute
	session :=4
	pomodoro := NewPomodoro(workDuration, shortBreak, longBreak, session)

	if pomodoro.workDuration != workDuration {
		t.Errorf("Expected work duration to be %v, got %v", workDuration, pomodoro.workDuration)
	}
	if pomodoro.shortBreak != shortBreak {
		t.Errorf("Expected short break duration to be %v, got %v", shortBreak, pomodoro.shortBreak)
	}
	if pomodoro.longBreak != longBreak {
		t.Errorf("Expected long break duration to be %v, got %v", longBreak, pomodoro.longBreak)
	}
	if pomodoro.sessions != session {
		t.Errorf("Expected session to be %v, got %v", session, pomodoro.sessions)
	}
}
