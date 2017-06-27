package main 

import (
	"fmt"
	"strings"
	"regexp"
)

var pattern = regexp.MustCompile(`\$\w+`)

func expand(s string, f func(string) string) string {
	return pattern.ReplaceAllStringFunc(s, func(s string) string { return f(s[1:]) })
}

// add 7 to all character of the string (similar to Caesar Cipher algorithm)
func toCaesar(s string) string {
	return strings.Map(func(r rune) rune {return r + 7}, s)
}

func main() {
	fmt.Printf(expand("This is a $test $Caesar $Encryption DONE!", toCaesar))
}