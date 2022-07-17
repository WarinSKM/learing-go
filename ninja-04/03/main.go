package main

import "fmt"

func main() {
	arr := []int{}

	for i := 42; i < 52; i++ {
		arr = append(arr, i)
	}

	fmt.Println(arr[:5])
	fmt.Println(arr[5:])
	fmt.Println(arr[2 : len(arr)-3])
	fmt.Println(arr[1 : len(arr)-4])
}
