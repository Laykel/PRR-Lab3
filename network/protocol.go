/*
Lab 2 - mutual exclusion
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
	RequestMessageType  = 0
	ReleaseMessageType  = 1
	SetValueMessageType = 2
)

// Read constants from parameters.json file
type Parameters struct {
	InitialPort    uint16 `json:"initial_port"`
	NbProcesses    uint8  `json:"nb_of_processes"`
	ProcessAddress string `json:"process_address1"`
}

var Params Parameters

// Message to request and release the critical section
type MessageCS struct {
	ReqType    uint8
	ProcessNbr uint8
	Timestamp  uint32
}

// Message to update the shared variable
type SetVariable struct {
	ReqType uint8
	Value   int32
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

// Decode bytes from MessageCS back to struct
func DecodeMessage(buffer []byte) MessageCS {
	message := MessageCS{}
	err := binary.Read(bytes.NewReader(buffer), binary.BigEndian, &message)
	if err != nil {
		log.Fatal(err)
	}

	return message
}

// Decode bytes from SetVariable back to struct
func DecodeSetVariable(buffer []byte) SetVariable {
	message := SetVariable{}
	err := binary.Read(bytes.NewReader(buffer), binary.BigEndian, &message)
	if err != nil {
		log.Fatal(err)
	}

	return message
}
