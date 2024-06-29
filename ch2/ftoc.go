package main

import "fmt"

func main() {
	fmt.Printf("Fahrenheit %g, Celsius %g\n", 32.0, ftoc(32.0))
	fmt.Printf("Fahrenheit %g, Celsius %g\n", 212.0, ftoc(212.0))
}

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}
