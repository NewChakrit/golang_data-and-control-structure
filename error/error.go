package error

import (
	"errors"
	"fmt"
)

func Error() {
	//result, err := divide(10, 0)
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//	return
	//}
	//fmt.Println("Result:", result)

	// ----------------------- //
	// Attempt to login with wrong credentials
	err := login("user", "pass")
	if err != nil {
		switch e := err.(type) {
		case *LogicError:
			// Custom error handing
			fmt.Println("Custom error occurred:", e)
		default:
			// Other types of errors
			fmt.Println("Generic error occurred:", e)

		}
		return
	}

	// Continue with the rest of the program if login is successful
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

// LogicError is a custom error type for login faillures
type LogicError struct {
	Username string
	Message  string
}

// Implement the Error() method to satisfy the error interface
func (e *LogicError) Error() string {
	return fmt.Sprintf("Login error for user '%s' : '%s' ", e.Username, e.Message)
}

// Simulate checking that attrempts a user login
func login(username, password string) error {
	//Simurate checkout username and password
	if username != "admin" || password != "password123" {
		return &LogicError{Username: username, Message: "invalid credential"}
	}
	// Login successful
	return nil
}
