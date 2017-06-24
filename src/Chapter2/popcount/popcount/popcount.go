package popcount

// array of 256 bytes
var pc [256]byte

// init fct cannot be called. init fct are automatically executed when the program starts
// pc[i] is the population count of i (in byte).
// For example if i = 27, then i(binary) = 11011 and pc[i] = 100 (4 in base 10)
func init() {
	// could have written
	// for i, _ := range pc
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}