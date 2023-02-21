package main

import (
	"fmt"
	"math"
)

func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

func main() {
	i := 5
	fmt.Println(Difference(i))
}

// another way
// func SquareOfSum(n int) int {
// 	sum := (n * (n + 1)) / 2
// 	return sum * sum
// }
// func SumOfSquares(n int) int {
// 	return (n * (n + 1) * (2*n + 1)) / 6
// }
// func Difference(n int) int {
// 	return SquareOfSum(n) - SumOfSquares(n)
// }
