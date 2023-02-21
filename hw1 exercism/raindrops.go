package main

import (
	"fmt"
	"strconv"
)

func Convert(number int) string {

	modes := [3]int{3, 5, 7}
	responses := [3]string{"Pling", "Plang", "Plong"}

	var answer string
	for index, value := range modes {
		if number%value == 0 {
			answer += responses[index]
		}
	}

	if answer == "" {
		answer = strconv.Itoa(number)
	}
	return answer
}

func main() {
	fmt.Println(Convert(250))
	fmt.Println(Convert(28))
	fmt.Println(Convert(30))
	fmt.Println(Convert(34))

}
