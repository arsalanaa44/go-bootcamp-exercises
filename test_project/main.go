package main

import "fmt"

func main() {

	i := 0
start:
	fmt.Println(i)
	i++
	if i <= 7 {
		goto start
	}
	fmt.Println("finished!")
}
