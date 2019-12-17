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
	"fmt"
	"log"
	"net"
	"time"
)

// Pass acknowledgment from listener to sender
var ack = make(chan ElectionMessage)

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

	for {
		// Read message from network
		message := ElectionMessage{}

		decoder := gob.NewDecoder(conn)
		err = decoder.Decode(&message)
		checkError(err)

		// Depending on the message type
		if message.MessageType != AcknowledgeMessageType {
			// Send message back to main routine
			req <- message

			// Send acknowledge message
			SendMessage(ElectionMessage{
                MessageType:      AcknowledgeMessageType,
                ProcessIdSender:  message.ProcessIdSender - 1,
            })
		} else {
			ack <- message
		}
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

	if message.MessageType != AcknowledgeMessageType {
		timeout := time.After(1 * time.Second)

		// Wait for acknowledgment
		select {
		case <-ack:
			return
		case <-timeout:
			fmt.Println("Timeout") // TODO Do something when we timeout
			return
		}
	}
}

// Simply crash if an error occurred
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// TODO doesn't work
func AreYouThere(address string) {
	for {
		// Connect to recipient's server
		conn, err := net.Dial("udp", address)

		if err == nil {
			conn.Close()
			break
		}
	}
}
