/*
Lab 3 - Chang and Roberts with failures
File: network/protocol.go
Authors: Jael Dubey, Luc Wachter
Go version: 1.13.4 (linux/amd64)

Describe networking values, messages structure for the protocol and provide
encoding and decoding functions for messages
*/
package network

import (
	"bytes"
	"encoding/binary"
	"log"
)

const (
	AnnounceMessageType    = 0
	ResultMessageType      = 1
	AcknowledgeMessageType = 2
)

type Announce struct {
	MessageType      uint8
	VisitedProcesses map[uint32]uint32 // Process number - aptitude
}

type Result struct {
	MessageType      uint8
	Elect            uint8
	VisitedProcesses []uint32 // Process number
}

type Acknowledge struct {
	MessageType uint8
}

// Encode given struct as big endian bytes and return bytes buffer
func Encode(message interface{}) []byte {
	buffer := &bytes.Buffer{}
	// Write struct's data as bytes
	err := binary.Write(buffer, binary.BigEndian, message)
	if err != nil {
		log.Fatal(err)
	}

	return buffer.Bytes()
}
