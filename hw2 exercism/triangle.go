package main

import (
	"fmt"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "not a triangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
)

func max(slc []float64) float64 {
	var m float64
	for i, e := range slc {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	var k Kind
	if a <= 0 || b <= 0 || c <= 0 || (2*(max([]float64{a, b, c})) >= (a + b + c)) {
		k = NaT
	} else {
		if a != b && b != c && c != a {
			k = Sca
		} else {
			sl := []float64{a, b, c, a, b}
			for i := 0; i < 3; i++ {
				if sl[i] == sl[i+1] {
					if sl[i+1] == sl[i+2] {
						k = Equ
					} else {
						k = Iso
					}
					break
				}
			}
		}
	}

	return k
}

// another great way
// import "math"
// Kind is the triangle kind
// type Kind int
// // Kind constants
// const (
// 	NaT Kind = iota // Not a triangle
// 	Equ             // Equilateral
// 	Iso             // Isosceles
// 	Sca             // Scalene
// )
// // KindFromSides tells if a triangle is equilateral, isosceles or scalene
// func KindFromSides(a, b, c float64) Kind {
// 	if math.IsNaN(a+b+c) || a+b <= c || a+c <= b || b+c <= a {
// 		return NaT
// 	}
// 	if a == b && a == c {
// 		return Equ
// 	}
// 	if a == b || a == c || b == c {
// 		return Iso
// 	}
// 	return Sca
// }

func main() {

	fmt.Println(
		KindFromSides(4, 3, 4),
		KindFromSides(1, 5, 4),
		KindFromSides(3, 3, 3),
		KindFromSides(5, 4, 3),
	)

}
