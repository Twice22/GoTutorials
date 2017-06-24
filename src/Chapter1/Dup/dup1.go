package main 

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create map with keys are strings and values are ints using built-in make fct
	counts := make(map[string]int)

	// create new variable that reads from the standard input
	input := bufio.NewScanner(os.Stdin)

	// read next line and removes the newline character from the end. It returns
	// true if there is a line and false otherwise
	for input.Scan() {
		// if a map doesn't yet contain a key, the value of counts[newkey]
		// is evaluated to the zero value of the corresping value type of the map (here 0
		// as ints are used as value for the counts map)

		// to quit if user doesn't write anything on the standard input
		if len(input.Text()) == 0 {
            break
        }

		// input.Text() retrieve the result of applying input.Scan()
		counts[input.Text()]++
	}

	for line, n := range counts {

		// as with for loop, if statement doesn't need braces around the condition
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}