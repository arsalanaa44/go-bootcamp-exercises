package main

import (
	"testing"
)

func TestIsLeapYear(t *testing.T) {

	for _, v := range testCases {
		result := IsLeapYear(v.year)
		if result != v.expected {
			t.Errorf(
				"\nexpected : %v ,"+
					"\nbut output is : %v",
				v.expected, result)
		}
	}
}

//go test .
