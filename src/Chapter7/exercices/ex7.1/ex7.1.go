package main 

import (
	"fmt"
	"bufio"
	"bytes"
)

type WordCounter int

// Write method of *ByteCounter type
func (c *WordCounter) Write(p []byte) (int, error) {

	words := 0 // count words

	input := bufio.NewScanner(bytes.NewReader(p))

	// set the split function to split words
	input.Split(bufio.ScanWords)

	for input.Scan() {
		words++
	}

	*c += WordCounter(words) // Convert int to ByteCounter (not a fct call)
	return words, nil
}


type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {

	lines := 0

	input := bufio.NewScanner(bytes.NewReader(p))

	for input.Scan() {
		lines++
	}

	*c += LineCounter(lines)
	return lines, nil
}

func main() {
	var w WordCounter
	w.Write([]byte("hello! There are exactly 9 words in this sentence"))
	fmt.Println(w) // "9"

	var l LineCounter
	l.Write([]byte(`There is
		exactly 3 lines
		in this example`))
	fmt.Println(l) // "3"
}

