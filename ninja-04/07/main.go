package main

import "fmt"

func main() {
	arr := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}

	fmt.Println(arr)
	for _, val := range arr {
		for _, value := range val {
			fmt.Println(value)
		}
	}
}
