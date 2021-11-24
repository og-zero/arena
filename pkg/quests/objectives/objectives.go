package objectives

import "fmt"

type (
	// Objective for quests and other events
	Objective interface {
		Start()
		Done()

		Description(...string) string
		Progress() float32
	}

	objective struct {
		id          uint64
		subject     string
		description string
		progress    float32
		started bool
	}
)

func New() *objective {
	return &objective{}
}

func (o *objective) Start() {
}

func (o *objective) Done() {
}

func (o *objective) Description(description ...string) string {
	if len(description) > 0 {
		for _, s := range description {
			o.description += fmt.Sprintf("%s\n", s)
		}
	} 

	return o.description
}

func (o *objective) Progress() float32 {
}

func (o objective) ID() uint64 {
	return o.id
}
