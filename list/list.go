package main

import "fmt"

func main() {
	var productNames [4]string
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	featurePrices := prices[1:3]
	// [1:3] meaning including index 1  and excluding index 3

	featurePrices1 := prices[:3]
	// [:3] meaning from starting to index 3

	featurePrices2 := prices[1:]
	// [1:] meaning from index 1 to end

	// In slicing no negetive numbers cannot be used

	featurePrices[0] = 199.99
	// can be reallocated
	fmt.Println(featurePrices, featurePrices1, featurePrices2)
	fmt.Println(len(featurePrices), cap(featurePrices))

}
