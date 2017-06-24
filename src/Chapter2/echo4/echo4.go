package main 

import (
	"fmt"
	"flag"
	"strings"
)

// create new flag variable of type bool, 3 arguments:
// - name of the flag (n)
// - default value (false)
// - message print if user provide an invalid argument or call -h (-help)
var n = flag.Bool("n", false, "omit trailing newline")

// flag.String takes a name, default value and a message and create a string var
var sep = flag.String("s", " ", "separator")

// Note:  sep and n are pointers to the flag var, so we need to use *sep or *n to
// access to the values
func main() {
	// update flag variables from their default values
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}