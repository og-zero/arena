package spells

type (
	Spell interface {
		Cast()
	}

	DamageSpell interface {
		Spell
		Damage() float64
	}

	HealingSpell interface {
		Spell
		Heal() float64
	}
)
