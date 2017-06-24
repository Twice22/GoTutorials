package popcount

func PopCount(x uint64) int {
	res := 0
	for i := uint(0); i < 64; i++ {
		res += int((x>>i) & 1)
	}
	return res;
}