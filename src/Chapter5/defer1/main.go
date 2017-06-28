package main 

import (
	"fmt"
)

func main() {
	f(3)
}

// when the program panic, all deferred fct
// are run in reverse order (stack principle)
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}