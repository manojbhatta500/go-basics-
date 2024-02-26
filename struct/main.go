package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello welcome to my struct program where")
	fmt.Println("we will calculate the area and parameter of rectangle")

	fmt.Println("***********************************************************")

	fmt.Println("enter the length of rectangle :")
	var length float64
	fmt.Scanln(&length)
	// fmt.Println("user length is ", length)

	fmt.Println("enter the width of rectangle :")
	var width float64
	fmt.Scanln(&width)
	fmt.Println("user length is ", width)

	userData := data{length: float64(length), width: float64(width)}

	calculateArea(userData)

	calculateParameter(userData)

}

type data struct {
	length float64
	width  float64
}

func calculateArea(manager data) {

	area := manager.length * manager.width
	fmt.Println("the area is ", area)

}

func calculateParameter(manager data) {

	parameter := manager.length * 4
	fmt.Println("parameter is ", parameter)

}
