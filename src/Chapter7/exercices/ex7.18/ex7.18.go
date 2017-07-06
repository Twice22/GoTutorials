package main 

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"bytes"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type xml.Name
	Attr []xml.Attr
	Children []Node
}

func exploreNode(n Node, w io.Writer, depth int) {
	switch n := n.(type) {
	case CharData:
		// depth*2 and "" are argument for %*s, to precise
		// that we indent by depth*2 with ""
		//fmt.Fprintf(w, "%*s%s\n", depth*2, "", string(n))
		fmt.Fprintf(w, "%*s<text>\n", depth*2, "")
	case *Element:
		// if n is an Element then we need to explore recursively
		// all the children
		fmt.Fprintf(w, "%*s<%s>\n", depth*2, "", n.Type.Local)
		for _, chid := range n.Children {
			exploreNode(chid, w, depth+1)
		}
	default:
		panic(fmt.Sprintf("got %T", n))
	}
}


func (el *Element) String() string {
	b := new(bytes.Buffer)
	exploreNode(el, b, 0)
	return b.String()
}

func parse(r io.Reader) (n Node, err error) {
	dec := xml.NewDecoder(os.Stdin)
	var stack []*Element // stack of elements

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		// type switch. switch if type (inside ())
		// if one of the case type (and use method expression accordingly)
		switch tok := tok.(type) {
		case xml.StartElement:
			element := &Element{tok.Name, tok.Attr, nil}

			// element can either be the root node, in this case stack = [element]
			// or it is the child of the previous element from the stack
			if len(stack) == 0 {
				n = element
				stack = append(stack, element)
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, element)
			}
			stack = append(stack, element)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			// if we reach to a text Node, then tok is a string containing a text
			// and we can link this text to the previous element of the stack (parent)
			// using CharData() to convert the string to CharData before pushing it to the stack
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, CharData(tok))
		}
	}
	return
}

// go build Chapter1/Fetch
// ./Fetch.exe https://www.w3.org/TR/2006/REC-xml11-20060816
func main() {
	node, err := parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing: %s", err)
		os.Exit(1)
	}
	fmt.Println(node)
}

