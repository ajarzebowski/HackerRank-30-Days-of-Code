package main

import (
	"fmt"

	"../day4/person"
)

type Person struct {
	age int
}

func main() {
	var number int

	fmt.Print("Give an age: ")
	_, err := fmt.Scanf("%d", &number)

	// Handle errors
	if err != nil {
		fmt.Println("Number Error!!!", err)
	}

	if number < 0 {
		number = 0
	}

	p := person.New(number)
	p.PersonAge()
}
