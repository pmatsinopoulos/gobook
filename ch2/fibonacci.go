package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n uint64

	n, err := strconv.ParseUint(os.Args[1], 10, 64)

	if err != nil {
		fmt.Println("Error %v", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", n)
	/*
		Given "n" we need to calculate the fibonacci number at position "n"
		if n = 0, then 0
		if n = 1, then 1
		if n = 2, then 1
		if n = 3, then 2
		if n = 4, then 3
		if n = 5, then 5
		if n = 6, then 8
		if n = 7, then 13
		if n = 8, then 21
		if n = 9, then 34
		...
		if n, then fib(n) <= fib(n-2) + fib(n-1), but this is recursive definition.
	*/
	if n == 0 {
		fmt.Println("Result = 0")
		os.Exit(0)
	}

	if n == 1 {
		fmt.Println("Result = 1")
		os.Exit(0)
	}

	n_2 := uint64(0)
	n_1 := uint64(1)
	var result uint64 = 0
	for i := uint64(2); i <= n; i++ {
		result = n_2 + n_1
		n_2, n_1 = n_1, result
	}

	fmt.Println("Result = ", result)
}
