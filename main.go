package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"log"
)

var host = flag.String("host", "127.0.0.1", "host")
var port = flag.String("port", "9090", "port")

func main() {
	flag.Parse()
	var listener net.Listener
	var err error
	listener, err = net.Listen("tcp", *host + ":" + *port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Listening on " + *host + ":" + *port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		var b = make([]byte, 20480)
		bytesRead, err := conn.Read(b)
		if err != nil {
			break
		}
		log.Printf("got: %s\n", string(b[:bytesRead]))
		var resp = ""
		conn.Write([]byte(resp))
	}

	log.Println("client disconnected")
}
