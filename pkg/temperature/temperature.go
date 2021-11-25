package temperature

const (
	defaultTemperature float32 = 0

	Celsius = iota
	Farenheit
)

type (
	Temperature interface {
		C() float64 // return temperature in Celsius
		F() float64
	}

	celsius     float32
	farenheit   float32
	temperature float32
)

func New() temperature {
	return temperature(defaultTemperature)
}

func (t temperature) toFarenheit() farenheit {
	// Formula: (0°C × 9/5) + 32 = 32°F
	return (farenheit(t) * 9 / 5) + 32
}

func (t temperature) toCelsius() celsius {
	// Formula: (0°C × 9/5) + 32 = 32°F
	return (celsius(t) - 32) * 5 / 9
}
