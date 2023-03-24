package main

import (
	"fmt"
	"strconv"
	"test_project/textcolor"
)

func main() {

	colorReset := "\033[0m"

	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"

	fmt.Println(string(textcolor.ColorRed) + "test")
	fmt.Println("another test")
	fmt.Println(string(colorGreen), "test")
	fmt.Println(string(colorYellow), "test")
	fmt.Println(string(colorBlue), "test")
	fmt.Println(string(colorPurple), "test")
	fmt.Println(string(colorWhite), "test")
	fmt.Println(string(colorCyan), "test", string(colorReset))
	fmt.Println("next")

	for i := 0; i < 50; i++ {

		fmt.Println(i, string("\033["+strconv.Itoa(i)+"m"), "color")
	}
}
