package main 

import (
	"fmt"
	"bytes"
	"os"
	"strings"
)

func divide(s string) (string, string) {
	if dot := strings.LastIndex(s, "."); dot > 0 {
		return s[:dot], s[dot:]
	}
	return s, ""
}

func comma(s string) string {
	var buf bytes.Buffer
	intpart, floatpart := divide(s)
	n := len(intpart)
	if n <= 3 {
		return s
	}
	for i, v := range intpart {
		if i > 0 && (n-i) % 3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(v))
	}

	// add float part
	fmt.Fprintf(&buf, "%s", string(floatpart))

	return buf.String()
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}