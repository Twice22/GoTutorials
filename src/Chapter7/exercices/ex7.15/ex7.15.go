package main 

import (
	"fmt"
	"os"
	"bytes"

	"Chapter7/exercices/ex7.15/eval"
)

func readInput() string {
	if len(os.Args) < 1 {
		return ""
	}

	var buf bytes.Buffer
	for _, s := range os.Args[1:] {
		buf.WriteString(s)
	}
	return buf.String()
}

// go run ex7.15.go "pow((3+x),2)*A/pi"
// go run ex7.15.go 'min(5, n!/((n-m)! * m!))'
// go run ex7.15.go 'min(5, 6)'
// Note: we must use a single quote ' if we use '!' as '!' is a character used
// by bash to print the history.
func main() {
	mathExpr := readInput()

	expr, err := eval.Parse(mathExpr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}



	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	myEnv := make(eval.Env)

	for v := range vars {
		fmt.Printf("Value of %s: ", string(v))
		var fv float64
		_, err := fmt.Scanf("%f", &fv)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Value of %s is incorrect:%f", string(v), fv)
		} else {
			myEnv[v] = fv
		}
	}

	fmt.Printf("Result:\t %.6g\n", expr.Eval(myEnv))
}