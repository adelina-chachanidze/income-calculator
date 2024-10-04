package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	//fmt.Print("Enter revenue: ")
	//fmt.Scan(&revenue)

	

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax rate: ")
	fmt.Scan(&taxRate)

	var beforeTax = revenue - expenses
	var afterTax = beforeTax - (beforeTax * (taxRate / 100))
	var ratio = beforeTax / afterTax

	fmt.Println("Your EBT:", beforeTax)
	fmt.Println("Your income:", afterTax)
	fmt.Println("Your ratio:", ratio)
}

func withInput (infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
	


	
}


