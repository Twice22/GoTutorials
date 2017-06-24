package main 

import (
	"fmt"
	"os"
)

func main() {
	// initialize variables s and sep to empty strings
	// We could have written
	// var s string
	// var s = ""
	// var s string = ""
	s, sep := "", ""

	// we use a range loop here. range argument retunr both the index and
	// the corresponding value. As we don't need the index we've just used
	// _ (we cannot used any other name like "temp" as unussed local variables
	// are not allowed in Go)
	for _, arg := range os.Arg[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}