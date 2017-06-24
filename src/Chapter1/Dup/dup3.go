package main 

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


// dup3 only read from files
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {

		// return a byte slice that must be converted into a string
		// so it can be split by strings.Split
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}