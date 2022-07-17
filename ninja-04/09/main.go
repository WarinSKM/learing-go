package main

import "fmt"

func main() {
	favorite := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for index, value := range favorite {
		fmt.Println(index)
		for _, v := range value {
			fmt.Printf("\t %s\n", v)
		}
	}

	favorite["todd"] = []string{"somthing", "someTime", "someOne"}

	for index, value := range favorite {
		fmt.Println(index)
		for _, v := range value {
			fmt.Printf("\t %s\n", v)
		}
	}

}
