package main 

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array
		// Grow by doubling, for amortized linear complexity
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}

		z = make([]int, zlen, zcap)

		// built-in fct that copy from one slice (x) to
		// another one (z). It returns the nb of elements copied
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

// "..." makes the fct variadic, that is to say it accepts
// any number of final arguments
func appendslice(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to expand the slice
		z = x[:zlen]
	} else {
		// There is insufficient space.
		// Grow by doubling, for amortized linear complexity
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		// create slice of len zlen and capacity zcap
		z = make([]int, zlen, zcap)
		copy (z, x)
	}
	copy(z[len(x):], y)
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	// how to supply a list of arg from a slice
	var u []int{1, 2, 3, 4, 5}
	u = appendslice(u, u...)

}