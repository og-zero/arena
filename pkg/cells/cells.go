package cells

import (
	characters "github.com/og-zero/arena/pkg/charachters"
)

type (
	cell struct {
		x           uint64
		y           uint64
		charachters []*characters.Charachter
	}

	Cell interface {
		X() uint64
		Y() uint64
		Characters() []*characters.Charachter
	}
)

func New() *cell {
	return &cell{}
}

func (c cell) X() uint64 {
	return c.x
}

func (c cell) Y() uint64 {
	return c.y
}

func (c cell) Charachters() []*characters.Charachter {
	return c.charachters
}
