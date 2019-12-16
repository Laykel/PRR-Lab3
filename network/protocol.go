/*
Lab 3 - Chang and Roberts with failures
File: network/protocol.go
Authors: Jael Dubey, Luc Wachter
Go version: 1.13.4 (linux/amd64)

Describe networking values, messages structure for the protocol and provide
encoding and decoding functions for messages
*/
package network

const (
	AnnounceMessageType    = 0
	ResultMessageType      = 1
	AcknowledgeMessageType = 2
)

// Message for an election
// Can either be an Announcement message, a Result message or an Ack
type ElectionMessage struct {
    MessageType      uint8
    Elect            uint8
    VisitedProcesses map[uint8]uint8 // Process number - aptitude
}

// TODO REMOVE ALL OTHER MESSAGE TYPE
type Announce struct {
	MessageType      uint8
	VisitedProcesses map[uint8]uint8 // Process number - aptitude
}

type Result struct {
	MessageType      uint8
	Elect            uint8
	VisitedProcesses map[uint8]bool // Process number
}

type Acknowledge struct {
	MessageType uint8
}

// Sends announce message to next server
// If no acknowledgment is received, ...
func SendAnnounce() {

}
