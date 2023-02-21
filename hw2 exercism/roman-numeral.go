package main

import (
	"errors"
	"fmt"
	"strconv"
)

var guide = [10][4]string{
	{"", "", "", ""},
	{"M", "C", "X", "I"},
	{"MM", "CC", "XX", "II"},
	{"MMM", "CCC", "XXX", "III"},
	{"", "CD", "XL", "IV"},
	{"", "D", "L", "V"},
	{"", "DC", "LX", "VI"},
	{"", "DCC", "LXX", "VII"},
	{"", "DCCC", "LXXX", "VIII"},
	{"", "CM", "XC", "IX"},
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("out_of_range")
	}
	inputToString := strconv.Itoa(input)
	l := len(inputToString)
	answer := ""
	for i := 0; i < l; i++ {
		answer = guide[input%10][3-i] + answer
		input /= 10
	}

	return answer, nil

}

func main() {

	fmt.Println(
		ToRomanNumeral(1345),
	)
	fmt.Println(
		ToRomanNumeral(2008),
	)
	fmt.Println(
		ToRomanNumeral(452),
	)
}

// 1	M	C	X	I
// 2	MM	CC	XX	II
// 3	MMM	CCC	XXX	III
// 4		CD	XL	IV
// 5		D	L	V
// 6		DC	LX	VI
// 7		DCC	LXX	VII
// 8		DCCC	LXXX	VIII
// 9		CM	XC	IX
