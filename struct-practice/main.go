package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(title, content)
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note Title:- ")

	if err != nil {
		fmt.Print(err)
		return "", "", err
	}
	content, err := getUserInput("Note Content:- ")

	if err != nil {
		fmt.Print(err)
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("Invalid input")
	}
	return value, nil
}
