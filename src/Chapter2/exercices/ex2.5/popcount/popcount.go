package popcount

func PopCount(x uint64) int {
	res := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		res++
	}
	return res;
}