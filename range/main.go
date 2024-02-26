package main

import "fmt"

func main() {

	var lists = []string{"pyton", "javaScript", "php", "dart", "Rust", "Go", "perl"}

	characterDisplay(lists)

}

func characterDisplay(list []string) {

	for i, e := range list {
		fmt.Println(i, ".", e)

		fmt.Println()

		for _, element := range e {
			fmt.Printf("%q\n", element)

		}

		fmt.Println()

	}

}
