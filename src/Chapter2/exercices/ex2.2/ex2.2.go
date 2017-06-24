package main

import (
	"fmt"
	"os"
	"strconv"
	"Chapter2/exercices/ex2.2/tempconv"
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

		fe := tempconv.Feet(t)
		m := tempconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n",
				fe, tempconv.FToM(fe), m, tempconv.MToF(m))

		p := tempconv.Pound(t)
		k := tempconv.Kg(t)
		fmt.Printf("%s = %s, %s = %s\n",
				p, tempconv.PToK(p), k, tempconv.KToP(k))
	}
}