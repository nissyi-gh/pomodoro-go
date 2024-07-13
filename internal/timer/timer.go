package timer

import "time"

type Timer struct {
	duration time.Duration
	ticker   *time.Ticker
	done     chan bool
}

func NewTimer(duration time.Duration) *Timer {
	return &Timer{
		duration: duration,
		done:     make(chan bool),
	}
}

func (t *Timer) Start() {
	t.ticker = time.NewTicker(time.Second)
	remaining := t.duration

	go func() {
		for {
			select {
			case <-t.ticker.C:
				remaining -= time.Second
				if remaining <= 0 {
					t.done <- true
					return
				}
			}
		}
	}()
}

func (t *Timer) Done() <-chan bool {
	return t.done
}

func (t *Timer) Stop() {
	if t.ticker != nil {
		t.ticker.Stop()
	}
}
