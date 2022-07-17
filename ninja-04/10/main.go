package main

import "fmt"

func main() {
	favorite := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println("FRIST ONE!!!!")
	printAllValueInMap(favorite)

	favorite["todd"] = []string{"somthing", "someTime", "someOne"}

	fmt.Println("SECOND ONE!!!!")
	printAllValueInMap(favorite)

	delete(favorite, "todd")
	fmt.Println("THRID ONE!!!!")
	printAllValueInMap(favorite)

}

func printAllValueInMap(val map[string][]string) {
	for index, value := range val {
		fmt.Println(index)
		for _, v := range value {
			fmt.Printf("\t %s\n", v)
		}
	}
}
