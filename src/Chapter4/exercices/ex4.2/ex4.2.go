package main 

import (
	"fmt"
	"flag"
	"os"
	"crypto/sha256"
	"crypto/sha512"
	"bytes"
)

// flag.String takes a name, default value and a message and create a string var
//var shatype = flag.String("s", " ", "384 for SHA384, 512 for SHA512")

const (
	UseSHA256 = iota
	UseSHA384
	UseSHA512
)


func main() {
	// update flag variables from their default values
	var method string

	myFlag := flag.NewFlagSet("", flag.ExitOnError)
	myFlag.StringVar(&method, "s", "256", "SHA method")
	myFlag.Parse(os.Args[1:])

	shaUse := UseSHA256

	// No break in golang for switch/case
	if len(os.Args) >= 2 {
		switch method {
		case "384":
			shaUse = UseSHA384
		case "512":
			shaUse = UseSHA512
		default:
			shaUse = UseSHA256
		}

		// word we want the SHA
		myWord := bytes.NewBufferString("")
		fmt.Fprintf(myWord, "%s", os.Args[len(os.Args)-1])
		
		s := myWord.String()

		if s != "" {
			switch shaUse {
			case UseSHA256:
				fmt.Printf("%x", sha256.Sum256([]byte(s)))
			case UseSHA384:
				fmt.Printf("%x", sha512.Sum384([]byte(s)))
			case UseSHA512:
				fmt.Printf("%x", sha512.Sum512([]byte(s)))
			}
		}
	}
}