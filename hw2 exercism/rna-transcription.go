package main

import (
	"fmt"
)

const (
	ddna = "ACGT"
	rrna = "UGCA"
)

func ToRNA(dna string) string {
	var myMap = map[rune]rune{
		'C': 'G',
		'G': 'C',
		'A': 'U',
		'T': 'A',
	}

	rna := ""
	for _, v := range dna {
		rna += string(myMap[v])
	}

	return rna
	// another way
	// var replacer = strings.NewReplacer(
	// 	"C", "G",
	// 	"G", "C",
	// 	"A", "U",
	// 	"T", "A",
	// )
	// return replacer.Replace(dna)
}

func main() {

	fmt.Println(ToRNA("ACGT"))
}
