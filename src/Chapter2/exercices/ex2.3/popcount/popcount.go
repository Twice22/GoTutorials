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
	var res byte

	// we need a uint because we shift operation only works with unsigned int
	var i uint
	for ;i < 8; i++ {
		res += pc[byte(x>>(i*8))]
	}
	return int(res)
}