package main

import (
	"errors"
	"fmt"
	"os"
)

const earningFile string = "earning.txt"

func writeToFile(ebt, eat, ratio float64) {
	earningString := fmt.Sprintf("Earning Before Tax:- %.1f\n Earning After Tax:- %.1f\n Ratio:- %.1f \n", ebt, eat, ratio)
	os.WriteFile(earningFile, []byte(earningString), 0644)
}
func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// fmt.Print("Revenue:- ")
	// outputPrint("Revenue:- ")
	// fmt.Scan(&revenue)
	revenue, err := getUserInput("Revenue:- ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Print("Expenses:- ")
	// outputPrint("Expenses:- ")
	// fmt.Scan(&expenses)
	expenses, err = getUserInput("Expenses:- ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Print("Tax Rate:- ")
	// outputPrint("Tax Rate:- ")
	// fmt.Scan(&taxRate)
	taxRate, err = getUserInput("Tax Rate:- ")
	if err != nil {
		fmt.Println(err)
		return
	}

	earningBeforeTax, earningAfterTax, ratio := calculateEarningValue(revenue, expenses, taxRate)

	fmt.Println("Earning Before Tax :- ", earningBeforeTax)
	fmt.Println("Earning After Tax :- ", earningAfterTax)
	fmt.Printf("Ratio:- %.1f", ratio)
	writeToFile(earningBeforeTax, earningAfterTax, ratio)

}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Enter a valid value")
	}
	return userInput, nil
}

func calculateEarningValue(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	eat := ebt * (1 - taxRate/100)
	ratio := ebt / eat

	return ebt, eat, ratio
}
