package main

import "fmt"

func main() {
	var arr [5]int
	for i := 1; i <= 5; i++ {
		arr[i-1] = i
	}

	for i := range arr {
		fmt.Println(i)
	}
	fmt.Printf("%T", arr)
}
