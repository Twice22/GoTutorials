package main 

import (
	"io"
	"log"
	"net"
	"os"
)

// go build netcat1.go
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	// read data from conn (src) and writes it to the standard
	// output (dst) until an end-of-file condition or an error occurs
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}