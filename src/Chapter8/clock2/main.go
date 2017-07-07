package main 

import (
	"io"
	"log"
	"net"
	"time"
)

// open on prompt
// go build main.go
// ./main.exe &
// go to ../netcat1/
// open an new prompt and type:
// ./netcat1.exe (you should see the time appearing in the prompt every second)
// in the first prompt type:
// ../netcat1/netcat1.exe (both prompt display the time every second. The program work concurrently!)
func main() {
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