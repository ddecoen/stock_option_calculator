package main

import "fmt"

func main() {
	var numOptions, strikePrice, salesPrice float64

	fmt.Print("Number of options: ")
	fmt.Scanln(&numOptions)

	fmt.Print("Strike price: $")
	fmt.Scanln(&strikePrice)

	fmt.Print("Sales price: $")
	fmt.Scanln(&salesPrice)

	gainPerOption := salesPrice - strikePrice
	totalGain := numOptions * gainPerOption

	fmt.Printf("\nGain per option: $%.2f\n", gainPerOption)
	fmt.Printf("Total potential gain: $%.2f\n", totalGain)
}