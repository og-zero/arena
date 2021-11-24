package weather

type (
	Weather interface {
		Start()
		Stop()
	}

	weather struct {
		running bool
		rain
		fog
		temperature
	}
)

func New() *Weather {
	return &weather{}
}

func (w *weather) Start() {
	// todo
}

func (w *weather) Stop() {
	// todo
}
