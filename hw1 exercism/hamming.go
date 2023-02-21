// https://exercism.org/tracks/go/exercises/hamming/edit
// package hamming
package main

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {

	var err error
	want := 0

	if len(a) != len(b) {
		err = errors.New("not equal length !!")
	}
	if err == nil {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				want++
			}
		}
	}
	return want, err
}
func main() {

	fmt.Println(Distance("aaa", "bbb"))
	fmt.Println(Distance("aaa", "aab"))
	fmt.Println(Distance("aaa1", "aa"))
	fmt.Println(Distance("", ""))
	fmt.Println(Distance("", "aab"))
	fmt.Println(Distance("aaa", "aba"))

}
