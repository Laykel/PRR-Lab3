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
	"strconv"
	"time"
)

// Pass acknowledgment from listener to sender
var ack = make(chan bool)

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
		switch message.MessageType {
		case AcknowledgeMessageType:
			ack <- true
		case EchoMessageType:
			// Send acknowledge message
			SendMeta(ElectionMessage{
				MessageType: AcknowledgeMessageType,
			}, message.ProcessIdSender)
		default:
			// Send message back to main routine
			req <- message

			// Send acknowledge message
			SendMeta(ElectionMessage{
				MessageType: AcknowledgeMessageType,
			}, message.ProcessIdSender)
		}
	}
}

// Listen TCP (used to ping)
func ListenTCP(address string, port int) {
	// Listen for incoming traffic
	listener, err := net.Listen("tcp", address+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		_, err := listener.Accept()
		if err != nil {
			log.Print("Error accepting connections: ", err)
			continue
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
	_ = encoder.Encode(message)

	timeout := time.After(2 * 1 * time.Second)

	if message.MessageType == EchoMessageType {
		select {
		case <-ack:
			EchoHaveResponse = true
			return
		case <-timeout:
			EchoHaveResponse = false
		}
	} else if message.MessageType != AcknowledgeMessageType {

		// Wait for acknowledgment
		select {
		case <-ack:
			return
		case <-timeout:
			fmt.Println("Timeout")
			SendMeta(message, (message.ProcessIdSender+1)%Params.NbProcesses)
			return
		}
	}
}


// Ping recipient
func AreYouThere(address string) {
	for {
		// Connect to recipient's server
		conn, err := net.Dial("tcp", address)

		if err == nil {
			conn.Close()
			break
		}
	}
}

// Simply crash if an error occurred
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
