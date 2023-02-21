package main

import "fmt"

// Proverb generate the relevant proverb for a given list
func Proverb(rhyme []string) []string {
	proverbialRhyme := make([]string, len(rhyme))
	for index := range rhyme {
		if index+1 == len(rhyme) {
			proverbialRhyme[index] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			proverbialRhyme[index] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[index], rhyme[index+1])
		}
	}
	return proverbialRhyme
}

// Proverb should have a comment documenting it.
func ProverbB(rhyme []string) []string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	var out = []string{}

	for i := 1; i < len(rhyme); i++ {
		out = append(out, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}
	if len(rhyme) > 0 {
		out = append(out, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}
	return out

}

func main() {
	rhyme := []string{"nail", "shoe", "horse", "rider", "message", "battle", "kingdom"}
	fmt.Println(Proverb(rhyme))
	fmt.Println(Proverb([]string{}))

}
