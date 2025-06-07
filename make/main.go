package main

import "fmt"

type floatMap map[string]float64

func main() {
	userNames := make([]string, 2, 5)
	// 5 is the maximum no that can be allocated to an array
	// 2 is a pre-allocated space to the array

	userNames = append(userNames, "Max")

	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	fmt.Println(courseRatings)
}
