package main 

import (
	"log"
	"net"
	"strings"
	"time"
	"fmt"
	"bufio"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		// use goroutine to allow for concurrent echo response
		go echo(c, input.Text(), 1*time.Second)
	}
	// Note: ignoring potential errors from input.Err()
	c.Close()
}

// go build
// ./reverb2.exe &
// in a cmd command prompt (not on github for windows prompt!!!)
// netcat2.exe
// your_text
func main() {
	l, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) //e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}