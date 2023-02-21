package main

import (
	"fmt"
)

func IsLeapYear(y int) bool {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	return (y%4 == 0) && (y%100 != 0 || y%400 == 0)
}

func main() {

	fmt.Println("hi")
}
