package main 

import (
	"fmt"
	"unicode"
)


func squashSpace(b []byte) []byte {
	res := b[:0]

	for i, c := range b {
		if unicode.IsSpace(rune(c)) {

			// trim if previous char was also a whitespace char
			// don't append anything to res
			if !(i > 0 && unicode.IsSpace(rune(b[i-1]))) {
				res = append(res, ' ')
			}
		} else {
			res = append(res, c)
		}
	}

	return res
}

func main() {
	b := []byte("abc\r  \n\rdef")

	fmt.Printf("%q\n", string(squashSpace(b)))
	fmt.Printf("%q\n", b)
}