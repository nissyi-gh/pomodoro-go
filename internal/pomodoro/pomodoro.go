package pomodoro

import (
	"fmt"
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
		fmt.Printf("Starting session %d\n", p.currentSession)
		p.timer = timer.NewTimer(p.workDuration)
	} else {
		fmt.Println("Taking a long break...")
		p.timer = timer.NewTimer(p.longBreak)
		p.currentSession = 0
	}
	p.timer.Start()
	p.displayTimeLeft()
}

func (p *Pomodoro) NextSession() {
	if p.currentSession == p.sessions {
		fmt.Println("Taking a long break...")
		p.timer = timer.NewTimer(p.longBreak)
		p.currentSession = 0
	} else {
		fmt.Println("Taking a short break...")
		p.timer = timer.NewTimer(p.shortBreak)
		p.currentSession++
	}
	p.timer.Start()
	p.displayTimeLeft()
}

func (p *Pomodoro) GetTimer() *timer.Timer {
	return p.timer
}

func (p *Pomodoro) displayTimeLeft() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
			case <-ticker.C:
				timeLeft := p.timer.TimeLeft()
				fmt.Printf("\r%s", formatDuration(timeLeft))
			case <-p.timer.Done():
				fmt.Println()
				return
		}
	}
}

func formatDuration(d time.Duration) string {
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
