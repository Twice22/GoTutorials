package main 

import "fmt"

// when these function are called with no arguments
// they will return 0
func max(vals ...int) int {
	if len(vals) != 0 {
		maxi := vals[0]
		for _, v := range vals {
			if v > maxi {
				maxi = v
			}
		}
		return maxi
	}
	return 0
}

func min(vals ...int) int {
	if len(vals) != 0 {
		mini := vals[0]
		for _, v := range vals {
			if v < mini {
				mini = v
			}
		}
		return mini
	}
	return 0
}

// min max that takes 1 argument at least
func max2(v int, vals ...int) int {
	if len(vals) != 0 {
		maxi := v
		for _, val := range vals {
			if val > maxi {
				maxi = val
			}
		}
		return maxi
	}
	return v
}

func min2(v int, vals ...int) int {
	if len(vals) != 0 {
		mini := v
		for _, val := range vals {
			if val < mini {
				mini = val
			}
		}
		return mini
	}
	return v
}

func main() {
	fmt.Println(max()) // "0"
	fmt.Println(max(8, 45, 2, -12)) // "45"
	fmt.Println(min(8, 45, 2, -12)) // "-12"

	fmt.Println(max2(8, 754, 65, -53)) // "754"
	fmt.Println(min2(8, 754, 65, -53)) // "-53"

	// fmt.Println(max2()) // error (need one argument at least)
}