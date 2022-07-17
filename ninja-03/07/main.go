package main

import "fmt"

func main() {
	if 2 != 2 {
		fmt.Println("2 == 2")
	} else if 222 == 22 {
		fmt.Println("222 == 22")
	} else {
		fmt.Println("else condition")
	}
}
