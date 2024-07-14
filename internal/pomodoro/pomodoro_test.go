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

func TestPomodoro_Start_FirstSession(t *testing.T) {
	workDuration := 5 * time.Millisecond
	shortBreak := 2 * time.Millisecond
	longBreak := 3 * time.Millisecond
	session := 2
	pomodoro := NewPomodoro(workDuration, shortBreak, longBreak, session)

	startTime := time.Now()
	pomodoro.Start()

	if pomodoro.currentSession != 1 {
		t.Errorf("Expected current session to be 1, got %v", pomodoro.currentSession)
	}

	<-pomodoro.timer.Done()
	elapsedTime := time.Since(startTime)

	if elapsedTime < workDuration {
		t.Errorf("Expected elapsed time to be at least %v, got %v", workDuration, elapsedTime)
	}
}

func TestPomodoro_NextSession_LastSession(t *testing.T) {
	workDuration := 5 * time.Millisecond
	shortBreak := 2 * time.Millisecond
	longBreak := 3 * time.Millisecond
	sessions := 2
	pomodoro := NewPomodoro(workDuration, shortBreak, longBreak, sessions)

	pomodoro.currentSession = sessions

	startTime := time.Now()
	pomodoro.NextSession()

	<-pomodoro.timer.Done()
	elapsedTime := time.Since(startTime)

	if elapsedTime < longBreak {
		t.Errorf("Expected elapsed time to be at least %v, got %v", longBreak, elapsedTime)
	}

	if pomodoro.currentSession != 0 {
		t.Errorf("Expected current session to be 0, got %d", pomodoro.currentSession)
	}
}

func TestPomodoro_NextSession_NotLastSession(t *testing.T) {
	workDuration := 5 * time.Millisecond
	shortBreak := 2 * time.Millisecond
	longBreak := 3 * time.Millisecond
	sessions := 2
	pomodoro := NewPomodoro(workDuration, shortBreak, longBreak, sessions)

	pomodoro.currentSession = sessions - 1

	startTime := time.Now()
	pomodoro.NextSession()

	<-pomodoro.timer.Done()
	elapsedTime := time.Since(startTime)

	if elapsedTime < shortBreak {
		t.Errorf("Expected elapsed time to be at least %v, got %v", shortBreak, elapsedTime)
	}

	if pomodoro.currentSession != sessions {
		t.Errorf("Expected current session to be %d, got %d", sessions, pomodoro.currentSession)
	}
}
