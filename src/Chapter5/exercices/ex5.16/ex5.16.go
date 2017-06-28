package main 

import (
	"fmt"
	"strings"
	"bytes"
)

// easy way using strings.Join
func stringJoin(sep string, slice ...string) string {
	return strings.Join(slice, sep)
}

// without using strings.Join
func stringJoin2(sep string, slice ...string) string {
	var buf bytes.Buffer
	length := len(slice)
	for i, s := range slice {
		buf.WriteString(s)
		if i != length - 1 {
			buf.WriteString(sep)
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(stringJoin(" ", "Hello,", "世界!", "How are you doing ?!", "Fine?"))
	fmt.Println(stringJoin2(" ", "Hello,", "世界!", "How are you doing ?!", "Fine?"))
}