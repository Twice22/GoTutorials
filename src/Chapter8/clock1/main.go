package main 

import (
	"io"
	"log"
	"net"
	"time"
)

// go build main.go
// ../netcat1/netcat1.exe

// if process is already running on a certain port just do (on windows)
// netstat -ano | findstr :yourPort
// taskkill /PID yourPID /F (NOTE: need to be in cmd and not in github for windows)

// Note: if we try to launch 2 clients, the 2nd client will have to wait until the first
// client is finished because the server is sequential. To avoid this we only need to add
// the go keyword to the call to handleConn (see clock2)
func main() {
	// Listen creates a net.Listener, an object that listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8000")
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
		handleConn(conn)
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