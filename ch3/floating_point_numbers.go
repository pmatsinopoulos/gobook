package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Floating point numbers starting...")

	fmt.Println("Print the powers of 'e'")

	for i := 0; i < 8; i++ {
		result := math.Exp(float64(i))

		fmt.Printf("power of e in %d = %8.3f\n", i, result)
	}

	fmt.Println("Print some special numbers:")

	var z float64

	fmt.Println(z, -z, 1/z, -1/z, z/z)

	fmt.Printf("Is z/z NaN %v\n", math.IsNaN(z/z))

	fmt.Printf("Printing a NaN using math.NaN %v\n", math.NaN())
}
