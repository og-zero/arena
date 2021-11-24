package objects

type (
	Object interface {
		X() int64
		Y() int64
		Type() uint64
	}
)
