package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue:- ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses:- ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate:- ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - taxRate/100)
	ratio := earningAfterTax / earningAfterTax

	fmt.Println("Earning Before Tax :- ", earningBeforeTax)
	fmt.Println("Earning After Tax :- ", earningAfterTax)
	fmt.Println("Ratio:- ", ratio)

}
