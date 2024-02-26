package main

import "fmt"

func main() {
	// declaring and using it
	var num int = 9
	var numptr *int = &num

	fmt.Println(num)

	fmt.Println(numptr)

	manager := pressed{hits: 1}
	fmt.Println("hits", manager.hits)

	countHits(&manager)

	//  this should return 2 because of
	fmt.Println("hits", manager.hits)

	//  now i will change the variable using the pointer lets see

	old := "hello"
	new := "world"

	fmt.Println("before")

	fmt.Println(old)

	fmt.Println(new)

	swapStrings(&old, &new)

	fmt.Println("after")

	fmt.Println(old)

	fmt.Println(new)

}

func swapStrings(old *string, new *string) {
	third := *old

	*old = *new

	*new = third

}

// now i will practise it using functions and struct

type pressed struct {
	hits int
}

func countHits(manager *pressed) {
	// the lexer will drefrence the pointer automatically  if its struct
	manager.hits = manager.hits + 1

}
