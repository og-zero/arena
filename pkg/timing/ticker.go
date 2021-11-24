package timing

import (
	"sync"
	"time"
)

type (
	ticker struct {
		time  time.Time
		tick  time.Duration
		async bool
		wait  *sync.WaitGroup
		do    []func()
	}
)

func NewTicker(tick time.Duration) *ticker {
	t := &ticker{
		time: time.Now(),
		tick: tick,
	}

	return t.run()
}

func (t *ticker) run() *ticker {
	for range time.NewTicker(t.tick).C {
		if t.async {
			for _, f := range t.do {
				go f()
			}
		} else {
			for _, f := range t.do {
				f()
			}
		}
	}

	return t
}
