package pomodoro

import (
	"time"

	"github.com/nissyi-gh/pomodoro-go/internal/timer"
)

type Pomodoro struct {
	workDuration   time.Duration
	shortBreak     time.Duration
	longBreak      time.Duration
	sessions       int
	currentSession int
	timer          *timer.Timer
}


func NewPomodoro(workDuration, shortBreak, longBreak time.Duration, sessions int) *Pomodoro {
	return &Pomodoro{
		workDuration:   workDuration,
		shortBreak:     shortBreak,
		longBreak:      longBreak,
		sessions:       sessions,
		currentSession: 0,
	}
}
