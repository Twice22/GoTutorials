package main 

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	input := bufio.NewScanner(file)

	// set the split fct to split words
	// this won't work for chinese text because there are
	// no spaces in chinese...
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		counts[word]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nword\tfreq\n")
	for w, c := range counts {
		fmt.Printf("%q:\t%d\n", w, c)
	}
}