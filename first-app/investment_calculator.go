package main

import (
	"fmt"
	"math"
)

const inflationRate = 8

func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRate float64 = 5.5
	var years int = 10

	// fmt.Print("Investment Amount:- ")
	outputText("Investment Amount:- ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Iexpected Return Rate:- ")
	outputText("Iexpected Return Rate:- ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years:- ")
	outputText("Years:- ")
	fmt.Scan(&years)

	futureValue, futureReactValue := calulateFutureValue(investmentAmount, expectedReturnRate, float64(years))

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureReactValue)

	// fmt.Println("Future Value: ",futureValue)
	// fmt.Printf("Future Value: %.1f\n Future Value (adjusted for Inflation): %.1f", futureValue, futureReactValue)

	fmt.Println(formattedFV, formattedRFV)

}

func outputText(text string) {
	fmt.Print(text)
}

func calulateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, float64(years))
	// return fv, rfv
	return
}

// func calulateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
// 	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv := fv / math.Pow(1+inflationRate/100, float64(years))
// 	return fv, rfv
// 	// return
// }
