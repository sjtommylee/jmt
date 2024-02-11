package main

import (
	"fmt"

	"github.com/0xsj/fn-go/pkg/either"
)

type User struct {
	Name string
	Age  int
}

// Validation function to check if the user's age is at least 18
func validateAge(user User) either.Either[string, User] {
	if user.Age >= 18 {
		return either.RightConstructor[user](user)
	} else {
		return either.LeftConstructor[string]("User's age must be at least 18")
	}
}

// Map function to apply validation function to a list of users
func validateUsers(users []User) []either.Either[string, User] {
	result := make([]either.Either[string, User], len(users))
	for i, u := range users {
		result[i] = validateAge(u)
	}
	return result
}

func Validate() {
	// Define a list of users
	users := []User{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 17},
		{Name: "Charlie", Age: 20},
	}

	// Validate the list of users
	validatedUsers := validateUsers(users)

	// Print validation results
	for i, result := range validatedUsers {
		if result.IsRight() {
			fmt.Printf("User %s is valid\n", users[i].Name)
		} else {
			fmt.Printf("User %s is invalid: %s\n", users[i].Name, result.Left)
		}
	}
}
