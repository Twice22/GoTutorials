package main

import (
	"fmt"
	"os"
	"strings"
	"net"
	"log"
	"bufio"
	"time"
)

type clock struct {
	name, host string
}

func splitNameHost(s string) (string, string) {
	res := strings.Split(s, "=")
	if len(res) != 2 {
		fmt.Errorf("Wrong format: %s\nExpected: NAME=HOST", s)
		return "", ""
	}
	return res[0], res[1]
}

// fonction call 2 times through go keyword (goroutine)
// for s.Scan() loop indefinitely since clock.go are sending the time every second on conn
func (c * clock) watch(conn net.Conn) {
	defer conn.Close()

	s := bufio.NewScanner(conn)
	for s.Scan() {
		fmt.Printf("\r%s: %s\n", c.name, s.Text())
	}
}


// ./clockWall.exe NY=localhost:8020 Tokyo=localhost:8030
func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: ./clockWall.exe NAME1=HOST1 [NAME2=HOST2...]")
		os.Exit(1)
	}

	clocks := make([]*clock, 0)

	for _, param := range os.Args[1:] {
		n, h := splitNameHost(string(param))
		clocks = append(clocks, &clock{n, h})
	}

	// same as netcat1 but for each host
	for _, c := range clocks {
		conn, err := net.Dial("tcp", c.host)
		if err != nil {
			log.Fatal(err)
		}

		go c.watch(conn)
	}

	// Sleep indefinitely while other goroutines from clock.go are writing on conn
	for {
		time.Sleep(time.Minute)
	}

}