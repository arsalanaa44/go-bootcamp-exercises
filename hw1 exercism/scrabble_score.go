package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Score(word string) int {
	Map_intitialize()
	word = strings.ToUpper(word)
	score := 0
	for _, s := range word {
		score += myMap[s]
	}

	return score
}

var myMap = map[rune]int{}

// var myMap = make(map[string]int)

func Map_intitialize() {
	for _, slice := range dataToStringSlice() {
		value := 0
		for i := len(slice) - 1; i >= 0; i-- {
			v, er := strconv.Atoi(string(slice[i]))
			if er != nil {
				break
			}
			value += v * int(math.Pow10(len(slice)-1-i))
		}
		for _, s := range slice[:len(slice)-1] {
			myMap[s] = value
		}
	}

}
func dataToStringSlice() []string {
	input := `  // A, E, I, O, U, L, N, R, S, T       1
				// D, G                               2
				// B, C, M, P                         3
				// F, H, V, W, Y                      4
				// K                                  5
				// J, X                               8
				// Q, Z                               10`
	for _, s := range "/ ,	" {
		input = strings.ReplaceAll(input, string(s), "")
	}

	return strings.Split(input, "\n")
}

func main() {

	fmt.Println(Score("cabbage"))
	fmt.Println(4)
}
