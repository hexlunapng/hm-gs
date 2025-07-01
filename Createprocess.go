package main

import (
	"bufio"
	"fmt"
	"net"
)


func main() {
	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Listening on port 7777...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept error: %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
}