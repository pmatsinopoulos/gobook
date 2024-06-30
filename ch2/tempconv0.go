package main

import "fmt"

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	Freezing      Celsius = 0
	BoilingC      Celsius = 100
)

func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Printf("Absolute Zero = %s, %g F\n", AbsoluteZeroC, CelsiusToFahrenheit(AbsoluteZeroC))
}
