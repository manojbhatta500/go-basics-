package main

import "fmt"

func main() {

	var actualroomdata = [4]room{
		{name: "dogada kedar", cleanstatus: true},

		{name: "vita kedar", cleanstatus: true},

		{name: "malikarjun", cleanstatus: true},
		{name: "bhagatpur", cleanstatus: false},
	}
	statusChecker(actualroomdata)
}

type room struct {
	name        string
	cleanstatus bool
}

func statusChecker(data [4]room) {
	for i := 0; i < len(data); i++ {
		if !data[i].cleanstatus {
			fmt.Println(data[i].name, " is not cleaned")
		} else {
			fmt.Println(data[i].name, "the room is cleaned")
		}
	}
}
