package tempconv

import "fmt"

func example1() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	// fmt.Printf("%g\n", boilingF-FreezingC) // compile error: type mismatch
}

func example2() {
	c := FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c) 			// "100°C"
	fmt.Printf("%g\n", c)   // "100°C"; does not call String
	fmt.Println(float64(c)) // "100°C"; does not call String
}