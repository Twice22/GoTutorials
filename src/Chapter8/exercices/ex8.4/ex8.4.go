package main 

import (
	"log"
	"net"
	"strings"
	"time"
	"fmt"
	"bufio"
	"sync"
)

func echo(c net.Conn, shout string, delay time.Duration, group sync.WaitGroup) {
	// Remove one goroutine from the WaitGroup
	defer group.Done() // similar to wg.Add(-1)
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup // number of working goroutines
	for input.Scan() {
		wg.Add(1) // add 1 before calling a goroutine
		// use goroutine to allow for concurrent echo response
		go echo(c, input.Text(), 1*time.Second, wg)
	}
	// Note: ignoring potential errors from input.Err()
	// closing write Half of the TCP connection when the number of active goroutines falls to zero
	go func() {
		wg.Wait()
		if tcpConn, ok := c.(*net.TCPConn); ok {
			tcpConn.CloseWrite()
		}	
	}()
}

// go build
// ./ex8.4.exe &
// in a cmd command prompt (not on github for windows prompt!!!)
// netcat3.exe
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