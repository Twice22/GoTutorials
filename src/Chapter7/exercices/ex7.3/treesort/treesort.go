package treesort 

import (
	"fmt"
	"bytes"
)

type Tree struct {
	value int
	left, right *Tree
}

// Sort sorts values in place
func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = Add(root, v)
	}
	AppendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice
func AppendValues(values []int, t *Tree) []int {
	if t != nil {
		values = AppendValues(values, t.left)
		values = append(values, t.value)
		values = AppendValues(values, t.right)
	}
	return values
}

func Add(t *Tree, v int) *Tree {
	if t == nil {
		t = new(Tree)
		t.value = v
		return t
	}
	if v < t.value {
		t.left = Add(t.left, v)
	} else {
		t.right = Add(t.right, v)
	}
	return t
}

// Answer to ex7.3
func (t *Tree) String() string {
	var buf bytes.Buffer

	 if t != nil {
	 	buf.WriteString(t.left.String())
	 	buf.WriteString(fmt.Sprintf("%v ", t.value))
	 	buf.WriteString(t.right.String())
	 }
	 return buf.String()
}