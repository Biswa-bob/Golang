package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (person User) OutputUserDetails() {
	//  (person user)  called a receiver argument or only receiver

	fmt.Println(person.firstName, person.lastName, person.birthdate)
}

func (person *User) ClearUserName() {
	person.firstName = ""
	person.lastName = ""

}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First Name, Last Name and Birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
