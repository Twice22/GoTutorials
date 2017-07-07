package main

import (
	"io"
	"log"
	"net"
	"time"
	"os"
	"flag"
	"fmt"
)


// TZ=US/Eastern ./clock.exe -p 8020 &
// TZ=Asia/Tokyo ./clock.exe -p 8030 &
func main() {
	var port string

	myFlag := flag.NewFlagSet("", flag.ExitOnError)
	myFlag.StringVar(&port, "p", "8000", "port number")
	myFlag.Parse(os.Args[1:])

	fmt.Printf("Listen on port: %s\n", port)
	listener, err := net.Listen("tcp", "localhost:" + port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Accept method blocks until an incoming connection is made
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}