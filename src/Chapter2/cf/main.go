// Cf converts its numeric argument to Celsius and Fahrenheit.
package main 

// import package name tempconv (tempconv is a directory containing .go files)
// use golang.org/x/tools/cmd/goimports tool to automatically handle imports
import (
	"fmt"
	"os"
	"strconv"
	"Chapter2/exercices/ex2.1/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}