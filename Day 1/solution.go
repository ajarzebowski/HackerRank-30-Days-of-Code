package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		integerVariable      = 5
		floatVariable        = 4.1
		stringVariable       = "Learning Golang "
		floatVariableStdin   float64
		integerVariableStdin int
		stringVariableStdin  string
	)

	// Cast input to integer
	fmt.Print("Type an integer: ")
	_, err := fmt.Scanf("%d", &integerVariableStdin)

	// Handle errors
	if err != nil {
		fmt.Println("Int Error!!!", err)
	}

	// Cast input to float
	fmt.Print("Type a float: ")
	_, err = fmt.Scanf("%f", &floatVariableStdin)

	// Handle errors
	if err != nil {
		fmt.Println("Float Error!!!", err)
	}

	// Printing variables from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type a string: ")
	stringVariableStdin, _ = reader.ReadString('\n')

	fmt.Println(integerVariable + integerVariableStdin)
	fmt.Println(floatVariable + floatVariableStdin)
	fmt.Println(stringVariable, stringVariableStdin)
}
