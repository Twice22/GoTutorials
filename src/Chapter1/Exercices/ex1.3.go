package main 

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func print_inef() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func print_ef() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	start := time.Now()
	print_inef()
	fmt.Printf("%d ns elapsed\n", time.Since(start).Nanoseconds())

	start2 := time.Now()
	print_ef()
	fmt.Printf("%d ns elapsed\n", time.Since(start2).Nanoseconds())

}