package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		mealPrice     float64
		tipPercentage int
		taxPercentage int
		totalCost     float64
	)
	fmt.Print("Meal price: ")
	_, err := fmt.Scanf("%f", &mealPrice)

	// Handle errors
	if err != nil {
		fmt.Println("Meal price Error!!!", err)
	}

	fmt.Print("Tip percentage: ")
	_, err = fmt.Scanf("%d", &tipPercentage)

	// Handle errors
	if err != nil {
		fmt.Println("Tip percentage Error!!!", err)
	}

	fmt.Print("Tax percentage: ")
	_, err = fmt.Scanf("%d", &taxPercentage)

	// Handle errors
	if err != nil {
		fmt.Println("Tax percentage Error!!!", err)
	}

	// Writing number with decimal places automatically cast it to float64 ex. 100.00

	tip := mealPrice * (float64(tipPercentage) / 100.00)
	tax := mealPrice * (float64(taxPercentage) / 100.00)

	totalCost = math.Round(mealPrice + tip + tax)

	// Print out solution always with two decimal places even despite the fact that the number is always being rounded
	fmt.Printf("%.2f\n", totalCost)
}
