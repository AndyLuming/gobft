package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	service := "127.0.0.2:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listen, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	//set 2 minutes timeout
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	// set maxium request length to 128B to prevent flood attack
	request := make([]byte, 128)
	defer conn.Close()
	for {
		readLen, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if readLen == 0 {
			// connection already closed by client
			break
		} else {
			msg := strings.TrimSpace(string(request[:readLen]))
			fmt.Println(msg)

			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

		//clear last read content
		request = make([]byte, 128)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
