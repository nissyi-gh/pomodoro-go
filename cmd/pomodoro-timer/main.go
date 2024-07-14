package main

import (
	"fmt"
	"time"

	"github.com/nissyi-gh/pomodoro-go/internal/pomodoro"
)

func main() {
	workDuration := 25 * time.Minute
	shortBreak := 5 * time.Minute
	longBreak := 15 * time.Minute
	sessions := 4

	p := pomodoro.NewPomodoro(workDuration, shortBreak, longBreak, sessions)

	for i := 0; i < sessions; i++ {
		p.Start()
		<-p.timer.Done()

		if i < sessions-1 {
			p.NextSession()
			<-p.timer.Done()
		}
	}

	p.NextSession()
	<-p.timer.Done()

	fmt.Println("Pomodoro completed!")
}
