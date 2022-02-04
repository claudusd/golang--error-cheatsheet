package main

import (
	"errors"
	"fmt"
)

type (
	UserNotFound struct {
		Cause    error
		Username string
	}

	DBError struct {
		Msg      string
		SQLError int
	}
)

var SQLError8 = DBError{SQLError: 8}
var SQLError6 = DBError{SQLError: 6}

func (e UserNotFound) Error() string {
	return fmt.Sprintf("user a %v not found", e.Username)
}

func (e UserNotFound) Unwrap() error {
	return e.Cause
}

func (e DBError) Error() string {
	return fmt.Sprintf("sql error %v, %v", e.SQLError, e.Msg)
}

func (e DBError) Is(target error) bool {
	t, ok := target.(DBError)
	if !ok {
		return false
	}
	return t.SQLError == e.SQLError
}

func getDB() error {
	return DBError{
		SQLError: 8,
		Msg:      "system error",
	}
}

func getUser(username string) error {
	err := getDB()
	return UserNotFound{
		Username: username,
		Cause:    err,
	}
}

func main() {
	err := getUser("claudusd")
	fmt.Printf("\tError: %v\n\n", err)
	if errors.As(err, &UserNotFound{}) {
		fmt.Printf("Wrap firts level: %v\n", err)
	}
	if errors.As(err, &DBError{}) {
		fmt.Printf("Wrap second level: %v\n", err)
	}

	if errors.Is(err, &UserNotFound{}) {
		fmt.Print("Error is UserNotFound\n")
	}

	if errors.Is(err, SQLError6) {
		fmt.Printf("SQL Error 6: %v\n", err)
	}
	if errors.Is(err, SQLError8) {
		fmt.Printf("SQL Error 8: %v\n", err)
	}
}
