package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func Valid(id string) bool {
	id_to_int := strings.Split(Reverse(strings.ReplaceAll(id, " ", "")), "")

	answer := false
	if len(id_to_int) < 2 {
		return false
	}

	mod := 0
	for i, v := range id_to_int {
		v_int, er := strconv.Atoi(v)
		if er != nil {
			return false
		}
		if i%2 == 1 {
			v_int *= 2
			if v_int > 9 {
				v_int += -9
			}
		}
		mod += v_int
	}
	if mod%10 == 0 {
		answer = true
	}
	return answer
}

func main() {

	fmt.Println(
		Valid("4539 3195 0343 6467"),
		Valid("8273 1232 7352 0569"),
		Valid("055b 444 285"),
		Valid("2"),
	)
}
