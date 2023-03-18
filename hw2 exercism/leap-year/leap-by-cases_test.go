package main

import (
	"testing"
)

func TestIsLeapYear(t *testing.T) {

	for _, v := range testCases {
		result := IsLeapYear(v.year)
		if result != v.expected {
			t.Errorf(
				"\nexpected : %v ,\nbut output is : %v", v.expected, result)
		}
	}
}

func TestIsLeapYear2(t *testing.T) {

	for _, v := range testCases {
		result := IsLeapYear2(v.year)
		if result != v.expected {
			t.Errorf(
				"\nexpected : %v ,\nbut output is : %v", v.expected, result)
		}
	}
}

var result bool

func BenchmarkIsLeapYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := IsLeapYear(i)
		result = res
	}
}

func BenchmarkIsLeapYear2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := IsLeapYear2(i)
		result = res
	}
}

//go test .
//go test -bench=. -count 2 -benchmem -benchtime=2s -cpu=2
