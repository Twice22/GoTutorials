package main 

import (
	"bufio"
	"fmt"
	"os"
)

// handle a list of file names
func main() {
	counts := make(map[string]string)
	dupFiles := make(map[string][]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, dupFiles)

			f.Close()
		}
	}

	// retrieve key value pair from counts using range loop
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, dupFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, dupFiles map[string][]string) {
	input := bufio.NewScanner(f)
	filename := f.Name()

	for input.Scan() {
		counts[input.Text()]++
	}
}