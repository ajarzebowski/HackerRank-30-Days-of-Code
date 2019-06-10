package main

import (
	// https://golang.org/pkg/bufio/
	"bufio"
	// https://golang.org/pkg/fmt/
	"fmt"
	// https://golang.org/pkg/os/
	"os"
)

func main() {
	// Printing variables from code
	var helloWorld = "Hello World!"
	shorthandHelloWorld := "Shorthand Hello World!"

	fmt.Println(helloWorld, "\n", shorthandHelloWorld)

	// Printing variables from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
