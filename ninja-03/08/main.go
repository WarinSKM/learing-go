package main

import "fmt"

func main() {
	switch {
	case 2 == 2:
		fmt.Println("2 == 2")
	case 2 == 3:
		fmt.Println("2 == 3")
	case 2 == 4:
		fmt.Println("2 == 4")
	case 3 == 3:
		fmt.Println("3 == 3")
	}
}
