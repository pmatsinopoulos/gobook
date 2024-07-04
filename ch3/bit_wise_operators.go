package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5 // 10 | 100000 => 100010 , dec => 34
	var y uint8 = 1<<1 | 1<<2 // 10 | 100 =>    000110, dec => 6
	var z = x & y             // 10, dec => 2
	var a = x | y             // 100110, dec => 38
	var b = x ^ y             // 100100, dec => 36
	var c = x &^ y            // 100000, dec => 32
	var d = x << 1            // 1000100, dec => 68
	var e = x >> 1            // 010001, dec => 17

	fmt.Printf("%08b, %8d\n", x, x)
	fmt.Printf("%08b, %8d\n", y, y)
	fmt.Printf("%08b, %8d\n", z, z)
	fmt.Printf("%08b, %8d\n", a, a)
	fmt.Printf("%08b, %8d\n", b, b)
	fmt.Printf("%08b, %8d\n", c, c)
	fmt.Printf("%08b, %8d\n", d, d)
	fmt.Printf("%08b, %8d\n", e, e)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // checks whether i bit is 1 or not
			fmt.Println(i) // it will print 1, and 5 since x has bit set to 1 for positions 1 and 5 from right to left
		}
	}

	fmt.Printf("%08b\n", x<<1)
}
