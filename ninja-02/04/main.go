package main

import "fmt"

func main() {
	x := 16
	fmt.Printf("%d \t %b \t %#x\n", x, x, x)
	bitShifted := x << 1
	fmt.Printf("%d \t %b \t %#x", bitShifted, bitShifted, bitShifted)
}
