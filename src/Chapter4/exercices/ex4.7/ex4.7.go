package main 

import (
	"fmt"
	"unicode/utf8"
)


// reverse a slice of ints in place
func rev(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}

func reverse(b []byte) []byte {
	for i := 0; i < len(b); {
		// reverse the rune
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i : i+size])
		i += size
	}

	// reverse the entire slice (as we already reverse
	// all the runes, the bytes are in the right order)
	rev(b)
	return b
}

func main() {
	b := []byte("This is 汉字。")
	fmt.Printf("%q\n", b)

	fmt.Printf("%q\n", reverse(b))
}