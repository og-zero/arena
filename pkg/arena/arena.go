package world

import (
	"fmt"

	"github.com/og-zero/arena/pkg/cells"
)

const (
	startX uint64 = 0
	startY uint64 = 0
)

type (
	world2D interface{}
	world3D interface{}

	world struct {
		x     uint64
		y     uint64
		cells map[uint64]map[uint64]*cells.Cell
	}

	charachter interface{}
)

func New() *world {
	return &world{}
}

func (a *world) Log() {
	fmt.Printf("Log:\n%+v\n", a)
}

func (a *world) Size(x, y uint64) *world {
	a.x = x
	a.y = y

	return a
}

func (a *world) init() error {
	switch {
	case a.x == 0:
		return fmt.Errorf("world 'x' can't be equal to 0")
	case a.y == 0:
		return fmt.Errorf("world 'y' can't be equal to 0")
	}

	for x := startX; x < a.x; x++ {
		a.cells[x] = make(map[uint64]*cells.Cell)
		for y := startY; y < a.y; y++ {
			a.cells[x] = make(map[uint64]*cells.Cell)
		}
	}

	return nil
}

func (a *world) Run() error {
	if err := a.init(); err != nil {
		return err
	}

	go a.run()
	return nil
}

func (a *world) run() {
}
