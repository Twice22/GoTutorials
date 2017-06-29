package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// ex6.5
// this expresion construct an unsigned interger of 0: uint(0)

// We then take the bitwise complement: ^uint(0)
// that is to say ^uint(0) is equivalent to m^uint(0) 
// where m = "all bits set to 1" (because uint(0) is an unsigned integer)

// so ^uint(0) = "all bits set to 1"

// we then shift right by 63
// - if we're operating on a 32-bit platform, then ^uint(0) = maximum value of uint32
//   and (^uint(0) >> 63) = 0 (because after 32 shift ^uint(0) is already 0 on a 32-bit platform)
// - if we're operating on a 64-bit platform, then ^uint(0) ) maximum value of uint64
//   and shifting right by 63 bit will lead to a value of 1

// then if we are on a 64-bit platform with shift 32 by 1 leading to bitSize = 64
// and if we are working on a 32-bit platform with shift 32 by 0 (we don't shift) leading to a value of 32
const bitSize = 32 << (^uint(0) >> 63)

// Has reports wether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	// for example if x = 227
	// then we need to set the 35th bit of words[3] to 1
	// because 3 * 64 + 35 = 227
	word, bit := x/bitSize, uint(x%bitSize)
	return word < len(s.words) && s.words[word]&(1 << bit) != 0
}

// Add adds the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/bitSize, uint(x%bitSize)

	// while s.words isn't big enough, increase the size of
	// words by one with 0 byte.
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	// once s.words is just the right size
	// set the bit-th bit of words[word] to 1
	// using xor operation
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		// we loop through t.words. If all the number
		// in tword CAN be in s.words (meaning of: i < len(s.words))
		// then simply xor all number of tword to those of s.words
		// indeed if a number is both in s and t then the xor operation
		// won't have any effect, yet if the number is only in t then
		// the xor operation will have for effect to add the number to
		// s since 0 ^ 1 = 1!
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
		// else, if the tword number cannot be in s.word (to big)
		// then simply append then to s.word (in right order due to the loop iteration)
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		// if word is 0 then no number in the
		// range [i*64, (i+1)*64] so continue 
		// to loop to find number of the set to display
		if word == 0 {
			continue
		}
		for j := 0; j < bitSize; j++ {
			// for each bit in [i*64, (i+1)*64], if
			// the i*64+j bit is set to 1 (we find a number of the set)
			if word&(1 << uint(j)) != 0 {

				// we use this condition to add a space between each number
				// except between '{' and 'first number', so that why we use '>'
				// and not '>='
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				// print the number
				fmt.Fprintf(&buf, "%d", bitSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// return the number of elements in a set
func (s *IntSet) Len() int {
	total := 0
	for _, b := range s.words {
		for b != 0 {
			b = b & (b - 1)
			total++
		}
	}
	return total
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/bitSize, uint(x%bitSize)

	// if the number CAN be in the set (there is enough space)
	if word <= len(s.words) {
		// then put a 0 using XOR operator in the right bit (1 ^ 1 = 0)
		s.words[word] ^= 1 << bit
	}
}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var sc IntSet
	for _, e := range s.words {
		sc.words = append(sc.words, e)
	}
	return &sc
}

// Answer to ex6.2
func (s *IntSet) AddAll(vals ...int) {
	// simply reuse Add method
	for _, v := range vals {
		s.Add(v)
	}
}

// Answer to ex6.3
// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	length := len(s.words)
	for i:= 0; i < length; i++ {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

// DifferenceWith sets s to the difference of s - t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &^= t.words[i]
		}
	}	
}

// DifferenceWith sets s to the difference of s - t
func (s *IntSet) SymetricDifference(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] ^= t.words[i]
		}
	}	
}

// Answer to ex6.4
func (s *IntSet) Elems() (elems []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitSize; j++ {
			if word&(1 << uint(j)) != 0 {
				elems = append(elems, bitSize*i+j)
			}
		}
	}
	// bare return
	return 
}