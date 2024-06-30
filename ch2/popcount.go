package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

// slice of 256 bytes,
// pc[0] is the population count of 0, i.e. 0
// pc[1] is the population count of 1, i.e. 1
// pc[2] is the population count of 2, i.e. 1
// pc[3] is the population count of 3, i.e. 2
// pc[4] is the population count of 4, i.e. 1
// pc[5] is the population count of 5, i.e. 2
// e.t.c.

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		// e.g. if i == 3, then 3 / 2 => 1, i&1, 3 & 1 => 1, hence 1 + 1 = 2
		// e.g. if i == 4, then 4 / 2 => 2 pc[2] = 1 + 0 => 1
		// e.g. if i == 5, then 5 / 2 => 2 pc[2] = 1 + 1 => 2
		// Note that i / 2 essentially mean shift right by 1 bit
	}
}

// PopCount()
func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func main() {
	intArg, _ := strconv.ParseUint(os.Args[1], 10, 64)

	fmt.Printf("Popcount for %d = %d\n", intArg, PopCount(intArg))
}
