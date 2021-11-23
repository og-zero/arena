package arena

import (
	"fmt"
)

const (
	startX uint64 = 0
	startY uint64 = 0
)

type (
	arena2D interface{}
	arena3D interface{}

	arena struct {
		x     uint64
		y     uint64
		cells map[uint64]map[uint64]*cell
	}

	charachter interface{}

	cell struct {
		x           uint64
		y           uint64
		charachters []*charachter
	}
)

func New() *arena {
	return &arena{}
}

func (a *arena) Log() {
	fmt.Printf("Log:\n%+v\n", a)
}

func (a *arena) Size(x, y uint64) *arena {
	a.x = x
	a.y = y

	return a
}

func (a *arena) init() error {
	switch {
	case a.x == 0:
		return fmt.Errorf("arena 'x' can't be equal to 0")
	case a.y == 0:
		return fmt.Errorf("arena 'y' can't be equal to 0")
	}

	for x := startX; x < a.x; x++ {
		a.cells[x] = make(map[uint64]struct{})
		for y := startY; y < a.y; y++ {
			a.cells[x] = make(map[uint64]struct{})
		}
	}
}

func (a *arena) Run() error {
	if err := a.init(); err != nil {
		return err
	}

	go a.run()
	return nil
}

func (a *arena) run() {
}
