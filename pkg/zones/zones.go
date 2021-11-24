package zones

type (
	Zone interface {
		ID() uint64
	}

	zone struct {
		id uint64
	}
)

func New() *zone {
	return &zone{}
}

func (z *zone) ID() uint64 {
	return z.id
}
