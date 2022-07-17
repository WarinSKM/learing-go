package main

import "fmt"

func main() {
	const favSport = "basketball"

	switch favSport {
	case "football":
		fmt.Println("yesss football!!!")
	case "basketball":
		fmt.Println("yesss basketball!!!")
	case "badminton":
		fmt.Println("yesss badminton!!!")
	default:
		fmt.Println("none of that")
	}
}
