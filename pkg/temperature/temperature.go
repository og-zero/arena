package temperature

type (
	Temperature interface {
		C() float64 // return temperature in Celsius
		F() float64
	}

	celsius   float32
	farenheit float32
)

const ()

func (c celsius) toFarenheit() farenheit {
	// Formula: (0°C × 9/5) + 32 = 32°F
	return (farenheit(c) * 9 / 5) + 32
}

func (f farenheit) toCelsius() celsius {
	// Formula: (0°C × 9/5) + 32 = 32°F
	return (celsius(f) - 32) * 5 / 9
}
