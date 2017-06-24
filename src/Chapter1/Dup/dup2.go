package main 

import (
	"bufio"
	"fmt"
	"os"
)

// handle a list of file names
func main() {
	// create map with keys are strings and values are ints using built-in make fct
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {

			// os.Open returns 2 values:
			// 1 - an open file (*os.File)
			// 2 - value of built-in error type (if err == nil, file successfully loaded)
			f, err := os.Open(arg)
			if err != nil {
				// %v is used to display any value type in a default format
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)

			// close the file and release any resources
			f.Close()
		}
	}

	// retrieve key value pair from counts using range loop
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// function may be declared in any particular order
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {

		// to quit if user doesn't write anything on the standard input
		if len(input.Text()) == 0 {
            break
        }

		counts[input.Text()]++
	}
}