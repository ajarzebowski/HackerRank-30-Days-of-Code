package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Give a number: ")
	_, err := fmt.Scanf("%d", &number)

	// Handle errors
	if err != nil {
		fmt.Println("Number Error!!!", err)
	}

	if number%2 == 0 {
		if number >= 2 && number <= 5 {
			fmt.Println("Not Weird. Between 2 and 5")
		} else if number >= 6 && number <= 20 {
			fmt.Println("Weird. Between 6 and 20")
		} else if number > 20 {
			fmt.Println("Not Weird. Above 20")
		} else {
			fmt.Println("Out of the range")
		}
	} else {
		fmt.Println("Weird. Odd number")
	}
}
