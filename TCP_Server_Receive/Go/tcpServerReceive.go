package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Start server...")

	// Listen on port 12345
	ln, _ := net.Listen("tcp", ":12345")

	// Accept connection
	conn, _ := ln.Accept()

	// Run loop forever (or until ctrl-c)
	for {
		// Get message, output
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
	}
}
