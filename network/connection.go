/*
Lab 3 - Chang and Roberts with failures
File: network/connection.go
Authors: Jael Dubey, Luc Wachter
Go version: 1.13.4 (linux/amd64)

Handles all reading and writing to the network.
*/
package network

import (
	"log"
	"net"
	"strings"
)

// Listen to UDP packets
func Listen(address string, port int, req chan string) {
	// Create UDP connection
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(address),
		Port: port,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)

		// Read packet from remote
		length, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}

		// Clean up message
		message := strings.TrimSpace(string(buffer[:length]))

		// Send message back to main routine
		req <- message
	}
}

// Send bytes to recipient
func Send(message []byte, address string, port string) {
	// Connect to recipient's server
	conn, err := net.Dial("udp", address+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Send encoded message
	_, err = conn.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
