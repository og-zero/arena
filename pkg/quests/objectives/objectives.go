package objectives

import "fmt"

type (
	// Objective for quests and other events
	Objective interface {
		Start()
		Done()

		Description(...string) string
		Progress(...float32) float32
	}

	objective struct {
		id          uint64
		subject     string
		description string
		progress    float32
		started     bool
		done        bool
	}
)

func New() *objective {
	return &objective{}
}

func (o *objective) Start() {
	o.started = true
	o.done = false
}

func (o *objective) Done() {
	o.started = false
	o.done = true
}

func (o *objective) Description(description ...string) string {
	if len(description) > 0 {
		for _, s := range description {
			o.description += fmt.Sprintf("%s\n", s)
		}
	}

	return o.description
}

func (o *objective) Progress(progress ...float32) float32 {
	if len(progress) > 0 {
		o.progress = progress[0]
	}

	return o.progress
}

func (o objective) ID() uint64 {
	return o.id
}
