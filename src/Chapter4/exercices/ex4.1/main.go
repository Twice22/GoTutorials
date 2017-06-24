package main 

import (
	"fmt"
	"os"
	"crypto/sha256"
)



func xor(a, b [32]byte) [32]byte {
	// we assume we use this function to compare sh256 only
	// so we can use [32]byte
	var dst [32]byte

	for i := 0; i < 32; i++ {
  		dst[i] = a[i] ^ b[i]
  	}
  	return dst
}

func PopCount(x uint64) int {
	res := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		res++
	}
	return res;
}

func nbSetBits(a [32]byte) int {
	res := 0

	for i := 0; i < 32; i++ {
		res += PopCount(uint64(a[i]))
	}
	return res
}

func main() {
	if len(os.Args) == 3 {
		c1 := sha256.Sum256([]byte(os.Args[1]))
		c2 := sha256.Sum256([]byte(os.Args[2]))
		diff := xor(c1, c2)
		fmt.Printf("%x\n%x\ndiff: %b\nnb bits that are differents: %d\n", c1, c2, diff, nbSetBits(diff))
	}
}