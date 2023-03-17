package main

import "testing"

func TestIsLeapYear(t *testing.T) {
	var testCases = []struct {
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
			expected:    true, //for test
		},
	}

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
