package main 

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// use Join function from strings package
	fmt.Println(strings.Join(os.Args[1:], " "))
}