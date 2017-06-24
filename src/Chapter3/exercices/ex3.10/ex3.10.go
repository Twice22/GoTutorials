package main 

import (
	"fmt"
	"bytes"
	"os"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	for i, v := range s {
		if i > 0 && (n-i) % 3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(v))
	}
	return buf.String()
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}