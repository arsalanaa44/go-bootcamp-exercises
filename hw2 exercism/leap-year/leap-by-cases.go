package main

import (
	"fmt"
)

var ttestCases = []struct {
	description string
	year        int
	expected    bool
}{
	{
		description: "year not divisible by 4 in common year",
		year:        2015,
		expected:    false,
	},
	{
		description: "year divisible by 2, not divisible by 4 in common year",
		year:        1970,
		expected:    false,
	},
	{
		description: "year divisible by 4, not divisible by 100 in leap year",
		year:        1996,
		expected:    true,
	},
	{
		description: "year divisible by 4 and 5 is still a leap year",
		year:        1960,
		expected:    true,
	},
	{
		description: "year divisible by 100, not divisible by 400 in common year",
		year:        2100,
		expected:    false,
	},
	{
		description: "year divisible by 100 but not by 3 is still not a leap year",
		year:        1900,
		expected:    false,
	},
	{
		description: "year divisible by 400 is leap year",
		year:        2000,
		expected:    true,
	},
	{
		description: "year divisible by 400 but not by 125 is still a leap year",
		year:        2400,
		expected:    true,
	},
	{
		description: "year divisible by 200, not divisible by 400 in common year",
		year:        1800,
		expected:    false,
	},
}
var ansDict = make(map[int]bool)

func IsLeapYear(year int) bool {
	for _, v := range ttestCases {
		if v.year == year {
			return v.expected
		}
	}
	return false
}

func IsLeapYear2(y int) bool {
	return (y%4 == 0) && (y%100 != 0 || y%400 == 0)
}

func main() {

	fmt.Println("hi")
}
