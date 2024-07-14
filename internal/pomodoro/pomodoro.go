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

func (p *Pomodoro) Start() {
	p.currentSession++

	if p.currentSession <= p.sessions {
		p.timer = timer.NewTimer(p.workDuration)
	} else {
		p.timer = timer.NewTimer(p.longBreak)
		p.currentSession = 0
	}
	p.timer.Start()
}

func (p *Pomodoro) NextSession() {
	if p.currentSession == p.sessions {
		p.timer = timer.NewTimer(p.longBreak)
		p.currentSession = 0
	} else {
		p.timer = timer.NewTimer(p.shortBreak)
		p.currentSession++
	}
	p.timer.Start()
}
