package arena

import "fmt"

type (
	arena2D interface{}
	arena3D interface{}

	arena struct {
		x uint64
		y uint64
		cells map[uint64]map[uint64]struct{}
	}
)

func New() (*arena) {
	return &arena{}
}

func (a *arena) Log() {
	fmt.Printf("Log:\n%+v\n", a)
}

func (a *arena) Size(x,y uint64) *arena {
	a.x = x
	a.y = y
	a.cells = make(map[uint64]map[uint64]struct{})
	return a
}