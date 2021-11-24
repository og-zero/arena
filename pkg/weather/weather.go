package weather

import "github.com/og-zero/arena/pkg/temperature"

type (
	Weather interface {
		Start()
		Stop()
	}

	weather struct {
		running bool
		rain
		fog
		temperature.Temperature
	}
)

func New() *weather {
	return &weather{}
}

func (w *weather) Start() {
	// todo
}

func (w *weather) Stop() {
	// todo
}
