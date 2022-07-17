package main

import "fmt"

func main() {
	yearOfBrith := 1999
	for {
		if yearOfBrith > 2022 {
			break
		}
		fmt.Println(yearOfBrith)
		yearOfBrith++
	}
}
