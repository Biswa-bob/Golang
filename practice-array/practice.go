package main

import "fmt"

func main() {
	hobbies := [3]string{"Coding", "Gyming", "Music Listening"}
	fmt.Println(hobbies)
	firstHobby := hobbies[:1]
	restHobby := hobbies[1:]
	fmt.Println(firstHobby)
	fmt.Println(restHobby)
	sliceHobbie := hobbies[:2]
	fmt.Println(sliceHobbie)
	hobbies1 := []string{}
	newArr := append(hobbies1, "Coding", "Gyming", "Music Listening")
	fmt.Print(newArr)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
