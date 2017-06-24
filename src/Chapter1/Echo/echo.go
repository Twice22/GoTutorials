package main

// import several packages using a list form
import (
	"fmt"
	"os"
)

func main() {
	// declare to variable (var) named s and sep of type string
	// if not explicitly initialized, var are initialized to their 0 values
	// for string: "", for number: 0
	var s, sep string

	// os.Args is a slice of string. os.Args[0] is the name of the command.
	// Here echo.go in go run echo.go test.
	for i := 1; i < len(os.Args); i++ {
		// concatenation
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// Note on for loop
/* for loop are the only loop statement in Go and the form is:

for initialization; condition; post {
	// statements
}

To write a "while" loop simply use:
for condition {
	
}

infinite loop:
for {
	
}

*/