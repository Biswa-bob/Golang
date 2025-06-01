package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------")
		// panic(err)
		// The Panic funciton helps to close the program with an error message
	}
	fmt.Println("Welcome to Go Bank!")

	for {

		presentOptions()

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)
		case 2:

			fmt.Println("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount must be greater than 0")
				return
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance Updated! New amount: ", accountBalance)
		case 3:
			fmt.Println("Withdrawal Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance && withdrawAmount <= 0 {
				fmt.Println("Invalid amount must be greater than 0 and less than the acount balance")
				// return
				continue
			}
			accountBalance -= withdrawAmount // accountBalance = accountBalance - withdrawAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance Updated! New amount: ", accountBalance)
		default:
			fmt.Println("Good Bye!")
			fmt.Println("Thank you for choosing Go Bank! Have a nice day")
			return
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is ", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Println("Withdrawal Amount: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount must be greater than 0")
		// 		return
		// 	}

		// 	accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
		// 	fmt.Println("Balance Updated! New amount: ", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Println("Your deposit: ")
		// 	var withdrawAmount float64
		// 	fmt.Scan(&withdrawAmount)
		// 	if withdrawAmount > accountBalance && withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount must be greater than 0 and less than the acount balance")
		// 		// return
		// 		continue
		// 	}
		// 	accountBalance -= withdrawAmount // accountBalance = accountBalance - withdrawAmount
		// 	fmt.Println("Balance Updated! New amount: ", accountBalance)
		// } else {
		// 	fmt.Println("Good Bye!")
		// 	break
		// }
	}

}
