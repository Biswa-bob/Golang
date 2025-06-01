package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFormFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}
	return balance, nil
}
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	// 0644 is a file permission meaning the owner will be read and write ehile others can only read it.
}

func main() {
	var accountBalance, err = getBalanceFormFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------")
		// panic(err)
		// The Panic funciton helps to close the program with an error message
	}
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
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
