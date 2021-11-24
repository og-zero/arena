package world

import (
	"fmt"

	"github.com/og-zero/arena/pkg/cells"
)

const (
	startX int64 = 0
	startY int64 = 0
	startZ int64 = 0
)

type (
	World interface {
		// Use
		New() World
		Run(World)

		// Init & Run
		init() error
		run()

		// Getters
		X() int64
		Y() int64
		// Setters
		setY(int64)
		setX(int64)
	}
	world struct {
		x     int64
		y     int64
		cells map[int64]map[int64]*cells.Cell
	}

	charachter interface{}
)

func New() world {
	return world{}
}

func (w *world) Log() {
	fmt.Printf("Log:\n%+v\n", w)
}

func (w *world) Size(x, y int64) *world {
	w.setX(x)
	w.setY(y)

	return w
}

func (w *world) setX(x int64) {
	w.x = x
}

func (w *world) setY(y int64) {
	w.y = y
}

func (w *world) X() int64 {
	return w.x
}

func (w *world) Y() int64 {
	return w.y
}

func (w *world) init() error {
	switch {
	case w.x == 0:
		return fmt.Errorf("world 'x' can't be equal to 0")
	case w.y == 0:
		return fmt.Errorf("world 'y' can't be equal to 0")
	}

	for x := startX; x < w.x; x++ {
		w.cells[x] = make(map[int64]*cells.Cell)
		for y := startY; y < w.y; y++ {
			w.cells[x] = make(map[int64]*cells.Cell)
		}
	}

	return nil
}

func Run(w World) error {
	if err := w.init(); err != nil {
		return err
	}

	go w.run()
	return nil
}

func (w *world) run() {
}
