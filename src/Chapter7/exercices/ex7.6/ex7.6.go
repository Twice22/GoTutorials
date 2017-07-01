package main 

import (
	"fmt"
	"flag"

	"Chapter7/exercices/ex7.6/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")


// go run ex7.6.go
// go run ex7.6.go -temp -18C
// go run ex7.6.go -temp 212°F
func main() {
	// parses the command-line flags from os.Args[1:]
	flag.Parse()

	// print the temperature using String() from celsiusFlag struct
	// i.e. using String() method from Celsius type embedded in celsiusFlag
	fmt.Println(*temp)
}