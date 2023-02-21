package main

import (
	"errors"
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
// /
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, v := range d {
		_, ok := h[v]
		if !ok {
			return h, errors.New("INVALID")
		}
		h[v]++

	}

	// another way
	for _, c := range d {
		if _, ok := h[c]; ok {
			h[c]++
		} else {
			return h, fmt.Errorf("invalid nucleotide")
		}
	}

	return h, nil
}

func main() {
	var d DNA
	d = "GATTACA"
	fmt.Println(d.Counts())
	d = "INVALID"
	fmt.Println(d.Counts())
}
