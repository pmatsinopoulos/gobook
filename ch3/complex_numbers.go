package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	fmt.Println("Complex numbers starting")

	var x complex128 = complex(1, 2) // 1 + 2i

	fmt.Printf("x = %v\n", x)

	var y complex128 = complex(3, 4) // 3 + 4i;

	fmt.Printf("y = %v\n", y)

	var xXy = x * y

	fmt.Printf("x*y = %v\n", xXy)

	fmt.Printf("real(x*y) = %v\n", real(xXy))
	fmt.Printf("imag(x*y) = %v\n", imag(xXy))

	fmt.Printf("1i * 1i = %v\n", 1i*1i)

	z := 3 + 5i
	w := 2 + 2i

	fmt.Printf("z = %v\n", z)
	fmt.Printf("w = %v\n", w)

	a := cmplx.Sqrt(-1)

	fmt.Printf("a = %v\n", a)
}
