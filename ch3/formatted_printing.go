package main

import "fmt"

func main() {
	o := 0666 // 6 + 6*8 + 6*8^2 = 438
	x := int64(0xdeadbeef)
	ascii := 'a'
	unicode := '⍄'
	newline := '\n'

	fmt.Printf("%d %[1]o %#[1]o\n", o)
	fmt.Printf("%d %[1]x, %#[1]x\n", x)
	fmt.Printf("%d %[1]c %[1]q", ascii)
	fmt.Println()
	fmt.Printf("%d %[1]c %[1]q", unicode)
	fmt.Println()
	fmt.Printf("%d %[1]c %[1]q", newline)
	fmt.Println()
}
