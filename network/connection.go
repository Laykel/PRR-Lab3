/*
Lab 3 - Chang and Roberts with failures
File: network/connection.go
Authors: Jael Dubey, Luc Wachter
Go version: 1.13.4 (linux/amd64)

Handles all reading and writing to the network.
*/
package network

import (
    "encoding/gob"
    "log"
	"net"
)

// Listen to UDP packets
func Listen(address string, port int, req chan ElectionMessage) {
	// Create UDP connection
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(address),
		Port: port,
	})
    checkError(err)
	defer conn.Close()

	// Initialize decoder
	gob.Register(ElectionMessage{})
	decoder := gob.NewDecoder(conn)

	for {
	    // Read message from network
	    message := ElectionMessage{}

        err = decoder.Decode(message)
        checkError(err)

		// Send message back to main routine
		req <- message
	}
}

// Send any struct to recipient as Gob
func SendGob(message ElectionMessage, address string, port int) {
	// Connect to recipient's server
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP(address),
		Port: port,
	})
    checkError(err)
	defer conn.Close()

	// Encode message as Gob and send it
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(message)
	checkError(err)

	// TODO wait for ACK
}

// Simply crash if an error occurred
func checkError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
