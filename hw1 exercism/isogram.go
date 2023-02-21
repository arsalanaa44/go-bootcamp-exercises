package main

import (
	"fmt"
	"strings"
)

func IsIsogram(word string) bool {

	word_split := strings.Split(strings.ToLower(word), "")
	m := make(map[string]int)
	for _, v := range word_split {
		_, ok := m[v]
		if ok {
			if v != " " && v != "-" {
				return false
			}
		} else {
			m[v] = 1
		}
	}
	return true
}

func main() {

	// fmt.Println(IsIsogram(("downstream")))
	// fmt.Println(IsIsogram(("six-year-old")))
	// fmt.Println(IsIsogram(("fucku")))
	// ss := []string{"s", "a"}``
	fmt.Println(IsIsogram("six-year-old"))
	fmt.Println(IsIsogram("sa"))
}
