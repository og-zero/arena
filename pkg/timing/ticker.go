package timing

type (
	ticker struct {
		time time.Time
		tick time.Duration
		wait *sync.WaitGroup
		do []func()
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
	for time.NewTicker(t.tick)

	return t
}
